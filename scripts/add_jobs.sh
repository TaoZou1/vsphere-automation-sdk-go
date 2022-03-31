#! /bin/bash
set -e

echo "Running add jobs script..."

mods=(
    "services/vsphere"
    "services/vmc"
    "services/vmc/draas"
    "services/vmc/autoscaler"
    "services/nsxt"
    "services/nsxt-gm"
    "services/nsxt-vmc-aws-integration"
    "services/nsxt-mp")

dir=$CI_PROJECT_DIR/.gitlab-ci

for module_path in ${mods[@]}
do
    module="${module_path////_}"
    echo "Adding job for module: $module"
    echo "Module path: $module_path"

    cat $dir/module-versioning-template.yml >> $dir/module-versioning.yml

    sed -i.bak "s#<<module>>#$module#" $dir/module-versioning.yml
    sed -i.bak "s#<<module_path>>#$module_path#" $dir/module-versioning.yml
done

rm -rf $dir/module-versioning.yml.bak
