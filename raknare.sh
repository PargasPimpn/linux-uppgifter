#!/bin/sh

tr ' ' '\n' | sort|uniq -c|sort | tac
