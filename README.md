# json2csv

Read JSON, write CSV

    usage: json2csv f1 f2:header2 f3 ... < some.json > some.csv

This program just reads stdin as a single JSON document, containing an array of
objects. It iterates over the objects, outputting CSV records for each object.
The fields, and order, are determined by the `f1 f2 ...` arguments.

example:

    json2csv id first_name:firstName last_name:lastName email < example/some.json > example/some.csv

(Example input and result in ./example)
