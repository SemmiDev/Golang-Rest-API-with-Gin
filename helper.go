package main

func transferModelToVo(model student) transformedStudent {
	var vo transformedStudent
	statusAktif := false
	if model.ActiveStatus == 1 {
		statusAktif = true
	} else {
		statusAktif = false
	}
	vo = transformedStudent{
		ID:           model.ID,
		Name:         model.Name,
		Address:      model.Address,
		PhoneNumber:  model.PhoneNumber,
		ActiveStatus: statusAktif,
		Class:        model.Class,
	}
	return vo
}

func transferVoToModel(vo transformedStudent) student {
	statusAktif := 0
	if vo.ActiveStatus == true {
		statusAktif = 1
	} else {
		statusAktif = 0
	}
	return student{
		Name:         vo.Name,
		Class:        vo.Class,
		PhoneNumber:  vo.PhoneNumber,
		ActiveStatus: statusAktif,
		Address:      vo.Address,
	}
}

func validatorCreated(vo transformedStudent) string {
	var kosong string = " Tidak Boleh Kosong"
	if vo.Name == "" {
		return "Name" + kosong
	}
	if vo.Address == "" {
		return "Address" + kosong
	}
	if vo.Class == "" {
		return "Class" + kosong
	}
	if vo.PhoneNumber == "" {
		return "No Hp" + kosong
	}
	return ""
}
