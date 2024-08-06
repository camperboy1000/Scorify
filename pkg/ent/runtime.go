// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/ent/check"
	"github.com/scorify/scorify/pkg/ent/checkconfig"
	"github.com/scorify/scorify/pkg/ent/inject"
	"github.com/scorify/scorify/pkg/ent/injectsubmission"
	"github.com/scorify/scorify/pkg/ent/minion"
	"github.com/scorify/scorify/pkg/ent/round"
	"github.com/scorify/scorify/pkg/ent/schema"
	"github.com/scorify/scorify/pkg/ent/scorecache"
	"github.com/scorify/scorify/pkg/ent/status"
	"github.com/scorify/scorify/pkg/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	checkMixin := schema.Check{}.Mixin()
	checkMixinFields0 := checkMixin[0].Fields()
	_ = checkMixinFields0
	checkFields := schema.Check{}.Fields()
	_ = checkFields
	// checkDescCreateTime is the schema descriptor for create_time field.
	checkDescCreateTime := checkMixinFields0[0].Descriptor()
	// check.DefaultCreateTime holds the default value on creation for the create_time field.
	check.DefaultCreateTime = checkDescCreateTime.Default.(func() time.Time)
	// checkDescUpdateTime is the schema descriptor for update_time field.
	checkDescUpdateTime := checkMixinFields0[1].Descriptor()
	// check.DefaultUpdateTime holds the default value on creation for the update_time field.
	check.DefaultUpdateTime = checkDescUpdateTime.Default.(func() time.Time)
	// check.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	check.UpdateDefaultUpdateTime = checkDescUpdateTime.UpdateDefault.(func() time.Time)
	// checkDescName is the schema descriptor for name field.
	checkDescName := checkFields[1].Descriptor()
	// check.NameValidator is a validator for the "name" field. It is called by the builders before save.
	check.NameValidator = checkDescName.Validators[0].(func(string) error)
	// checkDescSource is the schema descriptor for source field.
	checkDescSource := checkFields[2].Descriptor()
	// check.SourceValidator is a validator for the "source" field. It is called by the builders before save.
	check.SourceValidator = checkDescSource.Validators[0].(func(string) error)
	// checkDescWeight is the schema descriptor for weight field.
	checkDescWeight := checkFields[3].Descriptor()
	// check.WeightValidator is a validator for the "weight" field. It is called by the builders before save.
	check.WeightValidator = checkDescWeight.Validators[0].(func(int) error)
	// checkDescID is the schema descriptor for id field.
	checkDescID := checkFields[0].Descriptor()
	// check.DefaultID holds the default value on creation for the id field.
	check.DefaultID = checkDescID.Default.(func() uuid.UUID)
	checkconfigMixin := schema.CheckConfig{}.Mixin()
	checkconfigMixinFields0 := checkconfigMixin[0].Fields()
	_ = checkconfigMixinFields0
	checkconfigFields := schema.CheckConfig{}.Fields()
	_ = checkconfigFields
	// checkconfigDescCreateTime is the schema descriptor for create_time field.
	checkconfigDescCreateTime := checkconfigMixinFields0[0].Descriptor()
	// checkconfig.DefaultCreateTime holds the default value on creation for the create_time field.
	checkconfig.DefaultCreateTime = checkconfigDescCreateTime.Default.(func() time.Time)
	// checkconfigDescUpdateTime is the schema descriptor for update_time field.
	checkconfigDescUpdateTime := checkconfigMixinFields0[1].Descriptor()
	// checkconfig.DefaultUpdateTime holds the default value on creation for the update_time field.
	checkconfig.DefaultUpdateTime = checkconfigDescUpdateTime.Default.(func() time.Time)
	// checkconfig.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	checkconfig.UpdateDefaultUpdateTime = checkconfigDescUpdateTime.UpdateDefault.(func() time.Time)
	// checkconfigDescID is the schema descriptor for id field.
	checkconfigDescID := checkconfigFields[0].Descriptor()
	// checkconfig.DefaultID holds the default value on creation for the id field.
	checkconfig.DefaultID = checkconfigDescID.Default.(func() uuid.UUID)
	injectMixin := schema.Inject{}.Mixin()
	injectMixinFields0 := injectMixin[0].Fields()
	_ = injectMixinFields0
	injectFields := schema.Inject{}.Fields()
	_ = injectFields
	// injectDescCreateTime is the schema descriptor for create_time field.
	injectDescCreateTime := injectMixinFields0[0].Descriptor()
	// inject.DefaultCreateTime holds the default value on creation for the create_time field.
	inject.DefaultCreateTime = injectDescCreateTime.Default.(func() time.Time)
	// injectDescUpdateTime is the schema descriptor for update_time field.
	injectDescUpdateTime := injectMixinFields0[1].Descriptor()
	// inject.DefaultUpdateTime holds the default value on creation for the update_time field.
	inject.DefaultUpdateTime = injectDescUpdateTime.Default.(func() time.Time)
	// inject.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	inject.UpdateDefaultUpdateTime = injectDescUpdateTime.UpdateDefault.(func() time.Time)
	// injectDescTitle is the schema descriptor for title field.
	injectDescTitle := injectFields[1].Descriptor()
	// inject.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	inject.TitleValidator = injectDescTitle.Validators[0].(func(string) error)
	// injectDescID is the schema descriptor for id field.
	injectDescID := injectFields[0].Descriptor()
	// inject.DefaultID holds the default value on creation for the id field.
	inject.DefaultID = injectDescID.Default.(func() uuid.UUID)
	injectsubmissionMixin := schema.InjectSubmission{}.Mixin()
	injectsubmissionMixinFields0 := injectsubmissionMixin[0].Fields()
	_ = injectsubmissionMixinFields0
	injectsubmissionFields := schema.InjectSubmission{}.Fields()
	_ = injectsubmissionFields
	// injectsubmissionDescCreateTime is the schema descriptor for create_time field.
	injectsubmissionDescCreateTime := injectsubmissionMixinFields0[0].Descriptor()
	// injectsubmission.DefaultCreateTime holds the default value on creation for the create_time field.
	injectsubmission.DefaultCreateTime = injectsubmissionDescCreateTime.Default.(func() time.Time)
	// injectsubmissionDescUpdateTime is the schema descriptor for update_time field.
	injectsubmissionDescUpdateTime := injectsubmissionMixinFields0[1].Descriptor()
	// injectsubmission.DefaultUpdateTime holds the default value on creation for the update_time field.
	injectsubmission.DefaultUpdateTime = injectsubmissionDescUpdateTime.Default.(func() time.Time)
	// injectsubmission.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	injectsubmission.UpdateDefaultUpdateTime = injectsubmissionDescUpdateTime.UpdateDefault.(func() time.Time)
	// injectsubmissionDescGraded is the schema descriptor for graded field.
	injectsubmissionDescGraded := injectsubmissionFields[6].Descriptor()
	// injectsubmission.DefaultGraded holds the default value on creation for the graded field.
	injectsubmission.DefaultGraded = injectsubmissionDescGraded.Default.(bool)
	// injectsubmissionDescID is the schema descriptor for id field.
	injectsubmissionDescID := injectsubmissionFields[0].Descriptor()
	// injectsubmission.DefaultID holds the default value on creation for the id field.
	injectsubmission.DefaultID = injectsubmissionDescID.Default.(func() uuid.UUID)
	minionMixin := schema.Minion{}.Mixin()
	minionMixinFields0 := minionMixin[0].Fields()
	_ = minionMixinFields0
	minionFields := schema.Minion{}.Fields()
	_ = minionFields
	// minionDescCreateTime is the schema descriptor for create_time field.
	minionDescCreateTime := minionMixinFields0[0].Descriptor()
	// minion.DefaultCreateTime holds the default value on creation for the create_time field.
	minion.DefaultCreateTime = minionDescCreateTime.Default.(func() time.Time)
	// minionDescUpdateTime is the schema descriptor for update_time field.
	minionDescUpdateTime := minionMixinFields0[1].Descriptor()
	// minion.DefaultUpdateTime holds the default value on creation for the update_time field.
	minion.DefaultUpdateTime = minionDescUpdateTime.Default.(func() time.Time)
	// minion.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	minion.UpdateDefaultUpdateTime = minionDescUpdateTime.UpdateDefault.(func() time.Time)
	// minionDescDeactivated is the schema descriptor for deactivated field.
	minionDescDeactivated := minionFields[3].Descriptor()
	// minion.DefaultDeactivated holds the default value on creation for the deactivated field.
	minion.DefaultDeactivated = minionDescDeactivated.Default.(bool)
	roundMixin := schema.Round{}.Mixin()
	roundMixinFields0 := roundMixin[0].Fields()
	_ = roundMixinFields0
	roundFields := schema.Round{}.Fields()
	_ = roundFields
	// roundDescCreateTime is the schema descriptor for create_time field.
	roundDescCreateTime := roundMixinFields0[0].Descriptor()
	// round.DefaultCreateTime holds the default value on creation for the create_time field.
	round.DefaultCreateTime = roundDescCreateTime.Default.(func() time.Time)
	// roundDescUpdateTime is the schema descriptor for update_time field.
	roundDescUpdateTime := roundMixinFields0[1].Descriptor()
	// round.DefaultUpdateTime holds the default value on creation for the update_time field.
	round.DefaultUpdateTime = roundDescUpdateTime.Default.(func() time.Time)
	// round.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	round.UpdateDefaultUpdateTime = roundDescUpdateTime.UpdateDefault.(func() time.Time)
	// roundDescNumber is the schema descriptor for number field.
	roundDescNumber := roundFields[1].Descriptor()
	// round.NumberValidator is a validator for the "number" field. It is called by the builders before save.
	round.NumberValidator = roundDescNumber.Validators[0].(func(int) error)
	// roundDescComplete is the schema descriptor for complete field.
	roundDescComplete := roundFields[2].Descriptor()
	// round.DefaultComplete holds the default value on creation for the complete field.
	round.DefaultComplete = roundDescComplete.Default.(bool)
	// roundDescID is the schema descriptor for id field.
	roundDescID := roundFields[0].Descriptor()
	// round.DefaultID holds the default value on creation for the id field.
	round.DefaultID = roundDescID.Default.(func() uuid.UUID)
	scorecacheMixin := schema.ScoreCache{}.Mixin()
	scorecacheMixinFields0 := scorecacheMixin[0].Fields()
	_ = scorecacheMixinFields0
	scorecacheFields := schema.ScoreCache{}.Fields()
	_ = scorecacheFields
	// scorecacheDescCreateTime is the schema descriptor for create_time field.
	scorecacheDescCreateTime := scorecacheMixinFields0[0].Descriptor()
	// scorecache.DefaultCreateTime holds the default value on creation for the create_time field.
	scorecache.DefaultCreateTime = scorecacheDescCreateTime.Default.(func() time.Time)
	// scorecacheDescUpdateTime is the schema descriptor for update_time field.
	scorecacheDescUpdateTime := scorecacheMixinFields0[1].Descriptor()
	// scorecache.DefaultUpdateTime holds the default value on creation for the update_time field.
	scorecache.DefaultUpdateTime = scorecacheDescUpdateTime.Default.(func() time.Time)
	// scorecache.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	scorecache.UpdateDefaultUpdateTime = scorecacheDescUpdateTime.UpdateDefault.(func() time.Time)
	// scorecacheDescPoints is the schema descriptor for points field.
	scorecacheDescPoints := scorecacheFields[1].Descriptor()
	// scorecache.PointsValidator is a validator for the "points" field. It is called by the builders before save.
	scorecache.PointsValidator = scorecacheDescPoints.Validators[0].(func(int) error)
	// scorecacheDescID is the schema descriptor for id field.
	scorecacheDescID := scorecacheFields[0].Descriptor()
	// scorecache.DefaultID holds the default value on creation for the id field.
	scorecache.DefaultID = scorecacheDescID.Default.(func() uuid.UUID)
	statusMixin := schema.Status{}.Mixin()
	statusMixinFields0 := statusMixin[0].Fields()
	_ = statusMixinFields0
	statusFields := schema.Status{}.Fields()
	_ = statusFields
	// statusDescCreateTime is the schema descriptor for create_time field.
	statusDescCreateTime := statusMixinFields0[0].Descriptor()
	// status.DefaultCreateTime holds the default value on creation for the create_time field.
	status.DefaultCreateTime = statusDescCreateTime.Default.(func() time.Time)
	// statusDescUpdateTime is the schema descriptor for update_time field.
	statusDescUpdateTime := statusMixinFields0[1].Descriptor()
	// status.DefaultUpdateTime holds the default value on creation for the update_time field.
	status.DefaultUpdateTime = statusDescUpdateTime.Default.(func() time.Time)
	// status.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	status.UpdateDefaultUpdateTime = statusDescUpdateTime.UpdateDefault.(func() time.Time)
	// statusDescPoints is the schema descriptor for points field.
	statusDescPoints := statusFields[3].Descriptor()
	// status.PointsValidator is a validator for the "points" field. It is called by the builders before save.
	status.PointsValidator = statusDescPoints.Validators[0].(func(int) error)
	// statusDescID is the schema descriptor for id field.
	statusDescID := statusFields[0].Descriptor()
	// status.DefaultID holds the default value on creation for the id field.
	status.DefaultID = statusDescID.Default.(func() uuid.UUID)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescNumber is the schema descriptor for number field.
	userDescNumber := userFields[4].Descriptor()
	// user.NumberValidator is a validator for the "number" field. It is called by the builders before save.
	user.NumberValidator = userDescNumber.Validators[0].(func(int) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
