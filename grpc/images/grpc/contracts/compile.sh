#!/bin/bash
DSTDIR='../defs'
for pfile in $(find . -name '*.proto')
do
    reldir=$(echo $(dirname $pfile) | sed "s,^\.,$DSTDIR,g")
    mkdir -p $reldir
    fname=$(basename $pfile)
    fname=$(echo "${reldir}/${fname%.proto}.pb")
    protoc --descriptor_set_out=$fname $pfile
done
