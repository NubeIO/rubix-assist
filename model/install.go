package model

type Install struct {
	Name                            string `json:"name"`
	Version                         string `json:"version"`
	Source                          string `json:"source"`
	MoveExtractedFileToNameApp      bool   `json:"move_extracted_file_to_name_app"`
	MoveOneLevelInsideFileToOutside bool   `json:"move_one_level_inside_file_to_outside"`
	DeleteAppDataDir                bool   `json:"delete_app_data_dir"` // this will delete for example the db, plugins and config
}
