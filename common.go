package main

func validate(req *SetRequest) bool {
	return req.Entity != nil && req.Entity.Id > 0 && req.Entity.Name != ""
}
