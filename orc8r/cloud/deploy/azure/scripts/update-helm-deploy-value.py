#!/usr/bin/env python3
#
# Copyright (c) 2016-present, Facebook, Inc.
# All rights reserved.
#
# This source code is licensed under the BSD-style license found in the
# LICENSE file in the root directory of this source tree. An additional grant
# of patent rights can be found in the PATENTS file in the same directory.

# This script can be used to update a yaml file value. This is intended to be used at runtime
# by pipelines to set secret / deployment specific values

import argparse
import yaml

def main() -> None:
    args = _parse_args()
    inputFile = open(args.input)
    templateYaml = yaml.safe_load(inputFile)
    inputFile.close()
    parts = args.path.split('.')
    
    loopCount = len(parts)
    fileRef = templateYaml
    for i in range(0, loopCount):
        part = parts[i]
        print(part)
        if i == (loopCount - 1):
            # found the value to update
            fileRef[part] = args.value
        else:
            # navigating the yaml object tree
            fileRef = fileRef[part]

    outputFile = open(args.output, "w")
    yaml.dump(templateYaml, outputFile, default_flow_style = False, allow_unicode = True, encoding = None)
    outputFile.close()


def _parse_args() -> argparse.Namespace:
    """ Parse the command line args """
    parser = argparse.ArgumentParser(description='Orc8r Helm Input Generator')
    parser.add_argument('--output', '-o',
                        help="Path to output generated values file to", type=str)
    parser.add_argument('--input', '-i',
                        help="The input file", type=str)
    parser.add_argument('--path', '-p',
                        help="The path to the value to set", type=str)                        
    parser.add_argument('--value', '-v',
                        help="The value to set", type=str)

    args = parser.parse_args()
    return args


if __name__ == '__main__':
    main()