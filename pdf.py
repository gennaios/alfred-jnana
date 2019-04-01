#!/usr/bin/python
# -*- coding: utf-8 -*-

from __future__ import print_function

import json
import os
import sys

import fitz


def bookmarks_for_pdf(file_name):
    """
    Get bookmarks for PDF using PyMuPDF

    :param file_name: path to PDF file
    :return: parsed bookmarks
    """
    if not os.path.isfile(file_name) and not file_name.endswith('.pdf'):
        print("Expected PDF file.")
        sys.exit()

    return get_bookmarks(file_name)


def bookmarks_to_json(bookmarks):
    """
    Converts bookmarks from list of tuples to dictionary, and saves to JSON file

    :param bookmarks: bookmarks
    :return:
    """
    bookmark_list = []

    for bookmark in bookmarks:
        bookmark_dict = {
            "title": bookmark[0],
            "section": bookmark[1],
            "destination": bookmark[2]
        }
        bookmark_list.append(bookmark_dict)

    print(json.dumps(bookmark_list, indent=4))


def get_cover(pdf_file):
    """

    :param pdf_file:
    :return:
    """
    print("TODO %s" % pdf_file)


def get_bookmarks(pdf_file):
    """
    Get bookmarks from file and parse with parse_bookmarks()
    :param pdf_file: path to file
    :return: parsed bookmarks
    """
    doc = fitz.open(pdf_file)
    toc = doc.getToC(doc)

    bookmarks = parse_bookmarks(toc)

    if bookmarks:
        return bookmarks
    else:
        return None


def parse_bookmarks(raw_bookmarks):
    """
    Parse bookmarks from PyMuPDF, tuple of (title, section, destination)

    :param raw_bookmarks: list of lists from PyMuPDF
    :return: parsed bookmarks
    """
    sections = {}
    parsed_bookmarks = []

    for bookmark in raw_bookmarks:
        try:
            level = int(bookmark[0])
            title = bookmark[1].strip()
            destination = str(bookmark[2]).strip()

            section = ''
            sections[level] = title

            if level == 1:
                section = None
            else:
                current_level = level
                while current_level > 1:
                    if section == '':
                        section = sections[current_level - 1]
                    else:
                        section = (sections[current_level - 1] +
                                   ' > ' + section)

                    current_level -= 1

            if section is not None:
                section = section

            new_bookmark = (title, section, destination)
            parsed_bookmarks.append(new_bookmark)
        except Exception as e:
            # some bookmarks w/negative number or possibly others?
            print('error %s' % e.message)
            continue

    return parsed_bookmarks


def main():
    """

    :return:
    """
    args = sys.argv

    if not len(args) == 3:
        print("Expected three arguments: [command] [PDF file]")
    else:
        if args[1] == 'bookmarks':
            bookmarks = bookmarks_for_pdf(args[2])
            bookmarks_to_json(bookmarks)


if __name__ == '__main__':
    main()
