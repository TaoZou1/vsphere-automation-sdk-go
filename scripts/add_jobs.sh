#! /bin/bash
set -e

echo "Running add jobs script..."

mods=$(jq -r '.components | keys_unsorted | @tsv' $SDK_METADATA_FILE)
src_dir=$(jq -r '.go.srcDir' $SDK_METADATA_FILE)

dir=$CI_PROJECT_DIR/.gitlab-ci

for mod_name in ${mods[@]}
do
    module=$(jq -r '.components["'$mod_name'"].go.module' $SDK_METADATA_FILE)
    module_path="$src_dir/$module"
    module="${module_path////_}"
    echo "Adding job for module: $module"
    echo "Module path: $module_path"

    cat $dir/module-versioning-template.yml >> $dir/module-versioning.yml
    cat $dir/code-sync.yml >> $dir/module-versioning.yml

    sed -i.bak "s#<<module>>#$module#" $dir/module-versioning.yml
    sed -i.bak "s#<<module_path>>#$module_path#" $dir/module-versioning.yml
done

rm -rf $dir/module-versioning.yml.bak
