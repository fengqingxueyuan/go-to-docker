from __future__ import (
    absolute_import, print_function, division, unicode_literals
)

import logging
import re
import sys
from datetime import datetime

logging.basicConfig(level=logging.DEBUG)
logger = logging.getLogger(__name__)

CURRENT_YEAR = datetime.today().year

MIT_LICENSE_BLOB = """Copyright (c) %d Uber Technologies, Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.""" % CURRENT_YEAR

LICENSE_BLOB = """Copyright (c) %d Uber Technologies, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.""" % CURRENT_YEAR

MIT_LICENSE_BLOB_LINES_GO = [
    ('// ' + l).strip() + '\n' for l in MIT_LICENSE_BLOB.split('\n')
]

LICENSE_BLOB_LINES_GO = [
    ('// ' + l).strip() + '\n' for l in LICENSE_BLOB.split('\n')
]

COPYRIGHT_RE = re.compile(r'Copyright \(c\) (\d+)', re.I)


def update_go_license(name, force=False):
    with open(name) as f:
        orig_lines = list(f)
    lines = list(orig_lines)

    current_header = ''.join(lines[0:len(MIT_LICENSE_BLOB_LINES_GO)])
    mit_header = ''.join(MIT_LICENSE_BLOB_LINES_GO)
    if current_header == mit_header:
        lines = lines[len(MIT_LICENSE_BLOB_LINES_GO)+1:]

    found = False
    changed = False
    for i, line in enumerate(lines[:5]):
        m = COPYRIGHT_RE.search(line)
        if not m:
            continue

        found = True
        year = int(m.group(1))
        if year == CURRENT_YEAR:
            break

        new_line = COPYRIGHT_RE.sub('Copyright (c) %d' % CURRENT_YEAR, line)
        assert line != new_line, ('Could not change year in: %s' % line)
        lines[i] = new_line
        changed = True
        break

    if not found:
        if 'Code generated by' in lines[0]:
            lines[1:1] = ['\n'] + LICENSE_BLOB_LINES_GO
        else:
            lines[0:0] = LICENSE_BLOB_LINES_GO + ['\n']
        changed = True

    if changed:
        with open(name, 'w') as f:
            for line in lines:
                f.write(line)
        print(name)


def main():
    if len(sys.argv) == 1:
        print('USAGE: %s FILE ...' % sys.argv[0])
        sys.exit(1)

    for name in sys.argv[1:]:
        if name.endswith('.go'):
            try:
                update_go_license(name)
            except Exception as error:
                logger.error('Failed to process file %s', name)
                logger.exception(error)
                raise error
        else:
            raise NotImplementedError('Unsupported file type: %s' % name)


if __name__ == "__main__":
    main()
