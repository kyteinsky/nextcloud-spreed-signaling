#!/usr/bin/env python3
#
# Standalone signaling server for the Nextcloud Spreed app.
# Copyright (C) 2019 struktur AG
#
# @author Joachim Bauch <bauch@struktur.de>
#
# @license GNU AGPL version 3 or any later version
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.
#

try:
  # Fallback for Python2
  from cStringIO import StringIO
except ImportError:
  from io import StringIO
import csv
import subprocess
import sys

URL = 'https://raw.githubusercontent.com/datasets/country-codes/refs/heads/main/data/country-codes.csv'

def tostr(s):
  if isinstance(s, bytes) and not isinstance(s, str):
    s = s.decode('utf-8')
  return s

try:
  unicode
except NameError:
  # Python 3 files are returning bytes by default.
  def opentextfile(filename, mode):
    if 'b' in mode:
      mode = mode.replace('b', '')
    return open(filename, mode, encoding='utf-8')
else:
  def opentextfile(filename, mode):
    return open(filename, mode)

def country_fallback(entry):
  country = entry['ISO3166-1-Alpha-3']
  if not country:
    return ''

  if country == 'NAM':
    # Special case for Namibia
    return 'NA'
  else:
    return ''

def continent_fallback(entry):
  country = entry['ISO3166-1-Alpha-2']
  if not country:
    return ''

  region = entry['Region Name']
  assert region == 'Americas', 'Unknown entry: %r' % (entry)
  return 'NA'

def generate_map(filename):
  data = subprocess.check_output([
    'curl',
    '-L',
    URL,
  ])

  reader = csv.DictReader(StringIO(tostr(data)), delimiter=',')
  continents = {}
  for entry in reader:
    country = entry['ISO3166-1-Alpha-2'] or country_fallback(entry)
    continent = entry['Continent'] or continent_fallback(entry)
    if not country or not continent:
      continue

    continents.setdefault(country, []).append(continent)

  out = StringIO()
  out.write('package signaling\n')
  out.write('\n')
  out.write('// This file has been automatically generated, do not modify.\n')
  out.write('// Source: %s\n' % (URL))
  out.write('\n')
  out.write('var (\n')
  out.write('\tContinentMap = map[string][]string{\n')
  for country, continents in sorted(continents.items()):
    value = []
    for continent in continents:
      continent = '"%s"' % (continent)
      if not continent in value:
        value.append(continent)
    out.write('\t\t"%s": {%s},\n' % (country, ', '.join(value)))
  out.write('\t}\n')
  out.write(')\n')
  with opentextfile(filename, 'wb') as fp:
    fp.write(out.getvalue())

def main():
  if len(sys.argv) != 2:
    sys.stderr.write('USAGE: %s <filename>\n' % (sys.argv[0]))
    sys.exit(1)

  filename = sys.argv[1]
  generate_map(filename)

if __name__ == '__main__':
  main()
