#!/bin/bash
. ./config.sh
function exec()
{
if [[ $deploy == 'yes' ]]
then
echo "from functio"
fi
}
exec;
