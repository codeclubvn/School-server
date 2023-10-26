package mysql

import (
	"elearning/infra/mysql/model"

	"gorm.io/gorm"
)

func addUserTable(db *gorm.DB) error {
	var err error
	if err = db.AutoMigrate(&model.User{}); err != nil {
		return err
	}
	return nil
}

func addUserTokenTable(db *gorm.DB) error {
	var err error
	if err = db.AutoMigrate(&model.UserToken{}); err != nil {
		return err
	}
	return nil
}

func addCourseTable(db *gorm.DB) error {
	var err error
	if err = db.AutoMigrate((&model.Course{})); err != nil {
		return err
	}
	return nil
}

func addUnitTable(db *gorm.DB) error {
	var err error
	if err = db.AutoMigrate((&model.Unit{})); err != nil {
		return err
	}
	return nil
}

func addAssignmentTable(db *gorm.DB) error {
	var err error
	if err = db.AutoMigrate((&model.Assignment{})); err != nil {
		return err
	}
	return nil
}

func addRubricTable(db *gorm.DB) error {
	var err error
	if err = db.AutoMigrate((&model.Rubric{})); err != nil {
		return err
	}
	return nil
}

func addRubricCriteriaTable(db *gorm.DB) error {
	var err error
	if err = db.AutoMigrate((&model.RubricCriteria{})); err != nil {
		return err
	}
	return nil
}

func addRubricCriteriaRatingTable(db *gorm.DB) error {
	var err error
	if err = db.AutoMigrate((&model.RubricCriteriaRating{})); err != nil {
		return err
	}
	return nil
}

func addForeignKey(db *gorm.DB) error {
	var err error
	var query string 
	var result *gorm.DB
	if !db.Migrator().HasConstraint(&model.Unit{}, "fk_unit_course_id") {
		query = "ALTER TABLE school.school__unit ADD CONSTRAINT fk_unit_course_id FOREIGN KEY (course_id) REFERENCES school.school__course(id)"
		result = db.Exec(query)
		if err = result.Error; err != nil {
			return err
		}
	}
	if !db.Migrator().HasConstraint(&model.Assignment{}, "fk_assignment_unit_id") {
		query = "ALTER TABLE school.school__assignment ADD CONSTRAINT fk_assignment_unit_id FOREIGN KEY (unit_id) REFERENCES school.school__unit(id)"
		result = db.Exec(query)
		if err = result.Error; err != nil {
			return err
		}
	}
	if !db.Migrator().HasConstraint(&model.RubricCriteria{}, "fk_rubric_criteria_rubric_id") {
		query = "ALTER TABLE school.school__rubric_criteria ADD CONSTRAINT fk_rubric_criteria_rubric_id FOREIGN KEY (rubric_id) REFERENCES school.school__rubric(id)"
		result = db.Exec(query)
		if err = result.Error; err != nil {
			return err
		}
	}
	if !db.Migrator().HasConstraint(&model.RubricCriteriaRating{}, "fk_rubric_criteria_rating_rubric_criteria_id") {
		query = "ALTER TABLE school.school__rubric_criteria_rating ADD CONSTRAINT fk_rubric_criteria_rating_rubric_criteria_id FOREIGN KEY (rubric_criteria_id) REFERENCES school.school__rubric_criteria(id)"
		result = db.Exec(query)
		if err = result.Error; err != nil {
			return err
		}
	}
	return nil
}

