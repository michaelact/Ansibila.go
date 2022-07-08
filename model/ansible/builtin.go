package ansible

var (
	// Ansible Built-in Variable
	BuiltInVariable = []string{
		"ansible_check_mode", "ansible_config_file", "ansible_dependent_role_names", 
		"ansible_diff_mode", "ansible_forks", "ansible_inventory_sources", 
		"ansible_limit", "ansible_loop", "ansible_loop_var", "ansible_index_var", 
		"ansible_parent_role_names", "ansible_parent_role_paths", "ansible_play_batch", 
		"ansible_play_hosts", "ansible_play_hosts_all", "ansible_play_role_names", 
		"ansible_playbook_python", "ansible_role_names", "ansible_role_name", 
		"ansible_collection_name", "ansible_run_tags", "ansible_search_path", 
		"ansible_skip_tags", "ansible_verbosity", "ansible_version", "group_names", 
		"groups", "hostvars", "inventory_hostname", "inventory_hostname_short", 
		"inventory_dir", "inventory_file", "omit", "play_hosts", "ansible_play_name", 
		"playbook_dir", "role_name", "role_names", "role_path", "ansible" 
	}
)
