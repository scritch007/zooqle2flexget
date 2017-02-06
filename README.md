Convert the zooqle RSS format into something that flexget can understand.

Value will be exported into the guid ID field instead of the link field.

This means that in the flexget configuration you'll need to do

rss:
    - http://yourserver/XXXXX
    - guid