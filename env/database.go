package env

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// Database is a wrapper for my data storage tools.
// Unclear as to whether it is necessary, but I will keep it around
type Database struct {
	DB *sqlx.DB
}

func NewDatabase(driver string, dataSource string) *Database {
	db := &Database{}

	db.DB = sqlx.MustOpen(driver, dataSource)

	return db
}

//
// Errors
//

func (db *Database) Error(err error) error {
	if err == nil {
		return nil
	}

	// If it is a psql error, we want to convert it to a value we understand
	if pqError, ok := err.(*pq.Error); ok {

		switch pqError.Code {
		case psqlDuplicateColumn, psqlDuplicateObject, psqlUniqueViolation:
			return errors.New("Duplicate Entity")
		}

		// Some sort of DB errror happened, but we don't know what
		return errors.New("Database error occurred")
	}

	return err
}

//
// lib/pq errors
//

const (
	// Class 00 - Successful Completion
	psqlSuccessfulCompletion = "00000"
	// Class 01 - Warning
	psqlWarning                          = "01000"
	psqlDynamicResultSetsReturned        = "0100C"
	psqlImplicitZeroBitPadding           = "01008"
	psqlNullValueEliminatedInSetFunction = "01003"
	psqlPrivilegeNotGranted              = "01007"
	psqlPrivilegeNotRevoked              = "01006"
	psqlStringDataRightTruncationWarning = "01004"
	psqlDeprecatedFeature                = "01P01"
	// Class 02 - No Data (this is also a warning class per the SQL standard)
	psqlNoData                                = "02000"
	psqlNoAdditionalDynamicResultSetsReturned = "02001"
	// Class 03 - SQL Statement Not Yet Complete
	psqlSqlStatementNotYetComplete = "03000"
	// Class 08 - Connection Exception
	psqlConnectionException                           = "08000"
	psqlConnectionDoesNotExist                        = "08003"
	psqlConnectionFailure                             = "08006"
	psqlSqlclientUnableToEstablishSqlconnection       = "08001"
	psqlSqlserverRejectedEstablishmentOfSqlconnection = "08004"
	psqlTransactionResolutionUnknown                  = "08007"
	psqlProtocolViolation                             = "08P01"
	// Class 09 - Triggered Action Exception
	psqlTriggeredActionException = "09000"
	// Class 0A - Feature Not Supported
	psqlFeatureNotSupported = "0A000"
	// Class 0B - Invalid Transaction Initiation
	psqlInvalidTransactionInitiation = "0B000"
	// Class 0F - Locator Exception
	psqlLocatorException            = "0F000"
	psqlInvalidLocatorSpecification = "0F001"
	// Class 0L - Invalid Grantor
	psqlInvalidGrantor        = "0L000"
	psqlInvalidGrantOperation = "0LP01"
	// Class 0P - Invalid Role Specification
	psqlInvalidRoleSpecification = "0P000"
	// Class 0Z - Diagnostics Exception
	psqlDiagnosticsException                           = "0Z000"
	psqlStackedDiagnosticsAccessedWithoutActiveHandler = "0Z002"
	// Class 20 - Case Not Found
	psqlCaseNotFound = "20000"
	// Class 21 - Cardinality Violation
	psqlCardinalityViolation = "21000"
	// Class 22 - Data Exception
	psqlDataException                         = "22000"
	psqlArraySubscriptError                   = "2202E"
	psqlCharacterNotInRepertoire              = "22021"
	psqlDatetimeFieldOverflow                 = "22008"
	psqlDivisionByZero                        = "22012"
	psqlErrorInAssignment                     = "22005"
	psqlEscapeCharacterConflict               = "2200B"
	psqlIndicatorOverflow                     = "22022"
	psqlIntervalFieldOverflow                 = "22015"
	psqlInvalidArgumentForLogarithm           = "2201E"
	psqlInvalidArgumentForNtileFunction       = "22014"
	psqlInvalidArgumentForNthValueFunction    = "22016"
	psqlInvalidArgumentForPowerFunction       = "2201F"
	psqlInvalidArgumentForWidthBucketFunction = "2201G"
	psqlInvalidCharacterValueForCast          = "22018"
	psqlInvalidDatetimeFormat                 = "22007"
	psqlInvalidEscapeCharacter                = "22019"
	psqlInvalidEscapeOctet                    = "2200D"
	psqlInvalidEscapeSequence                 = "22025"
	psqlNonstandardUseOfEscapeCharacter       = "22P06"
	psqlInvalidIndicatorParameterValue        = "22010"
	psqlInvalidParameterValue                 = "22023"
	psqlInvalidRegularExpression              = "2201B"
	psqlInvalidRowCountInLimitClause          = "2201W"
	psqlInvalidRowCountInResultOffsetClause   = "2201X"
	psqlInvalidTimeZoneDisplacementValue      = "22009"
	psqlInvalidUseOfEscapeCharacter           = "2200C"
	psqlMostSpecificTypeMismatch              = "2200G"
	psqlNullValueNotAllowed                   = "22004"
	psqlNullValueNoIndicatorParameter         = "22002"
	psqlNumericValueOutOfRange                = "22003"
	psqlStringDataLengthMismatch              = "22026"
	psqlStringDataRightTruncationException    = "22001"
	psqlSubstringError                        = "22011"
	psqlTrimError                             = "22027"
	psqlUnterminatedCString                   = "22024"
	psqlZeroLengthCharacterString             = "2200F"
	psqlFloatingPointException                = "22P01"
	psqlInvalidTextRepresentation             = "22P02"
	psqlInvalidBinaryRepresentation           = "22P03"
	psqlBadCopyFileFormat                     = "22P04"
	psqlUntranslatableCharacter               = "22P05"
	psqlNotAnXmlDocument                      = "2200L"
	psqlInvalidXmlDocument                    = "2200M"
	psqlInvalidXmlContent                     = "2200N"
	psqlInvalidXmlComment                     = "2200S"
	psqlInvalidXmlProcessingInstruction       = "2200T"
	// Class 23 - Integrity Constraint Violation
	psqlIntegrityConstraintViolation = "23000"
	psqlRestrictViolation            = "23001"
	psqlNotNullViolation             = "23502"
	psqlForeignKeyViolation          = "23503"
	psqlUniqueViolation              = "23505"
	psqlCheckViolation               = "23514"
	psqlExclusionViolation           = "23P01"
	// Class 24 - Invalid Cursor State
	psqlInvalidCursorState = "24000"
	// Class 25 - Invalid Transaction State
	psqlInvalidTransactionState                         = "25000"
	psqlActiveSqlTransaction                            = "25001"
	psqlBranchTransactionAlreadyActive                  = "25002"
	psqlHeldCursorRequiresSameIsolationLevel            = "25008"
	psqlInappropriateAccessModeForBranchTransaction     = "25003"
	psqlInappropriateIsolationLevelForBranchTransaction = "25004"
	psqlNoActiveSqlTransactionForBranchTransaction      = "25005"
	psqlReadOnlySqlTransaction                          = "25006"
	psqlSchemaAndDataStatementMixingNotSupported        = "25007"
	psqlNoActiveSqlTransaction                          = "25P01"
	psqlInFailedSqlTransaction                          = "25P02"
	// Class 26 - Invalid SQL Statement Name
	psqlInvalidSqlStatementName = "26000"
	// Class 27 - Triggered Data Change Violation
	psqlTriggeredDataChangeViolation = "27000"
	// Class 28 - Invalid Authorization Specification
	psqlInvalidAuthorizationSpecification = "28000"
	psqlInvalidPassword                   = "28P01"
	// Class 2B - Dependent Privilege Descriptors Still Exist
	psqlDependentPrivilegeDescriptorsStillExist = "2B000"
	psqlDependentObjectsStillExist              = "2BP01"
	// Class 2D - Invalid Transaction Termination
	psqlInvalidTransactionTermination = "2D000"
	// Class 2F - SQL Routine Exception
	psqlSqlRoutineException               = "2F000"
	psqlFunctionExecutedNoReturnStatement = "2F005"
	psqlModifyingSqlDataNotPermitted      = "2F002"
	psqlProhibitedSqlStatementAttempted   = "2F003"
	psqlReadingSqlDataNotPermitted        = "2F004"
	// Class 34 - Invalid Cursor Name
	psqlInvalidCursorName = "34000"
	// Class 38 - External Routine Exception
	psqlExternalRoutineException                = "38000"
	psqlExternalContainingSqlNotPermitted       = "38001"
	psqlExternalModifyingSqlDataNotPermitted    = "38002"
	psqlExternalProhibitedSqlStatementAttempted = "38003"
	psqlExternalReadingSqlDataNotPermitted      = "38004"
	// Class 39 - External Routine Invocation Exception
	psqlExternalRoutineInvocationException = "39000"
	psqlInvalidSqlstateReturned            = "39001"
	psqlExternalNullValueNotAllowed        = "39004"
	psqlTriggerProtocolViolated            = "39P01"
	psqlSrfProtocolViolated                = "39P02"
	// Class 3B - Savepoint Exception
	psqlSavepointException            = "3B000"
	psqlInvalidSavepointSpecification = "3B001"
	// Class 3D - Invalid Catalog Name
	psqlInvalidCatalogName = "3D000"
	// Class 3F - Invalid Schema Name
	psqlInvalidSchemaName = "3F000"
	// Class 40 - Transaction Rollback
	psqlTransactionRollback                     = "40000"
	psqlTransactionIntegrityConstraintViolation = "40002"
	psqlSerializationFailure                    = "40001"
	psqlStatementCompletionUnknown              = "40003"
	psqlDeadlockDetected                        = "40P01"
	// Class 42 - Syntax Error or Access Rule Violation
	psqlSyntaxErrorOrAccessRuleViolation   = "42000"
	psqlSyntaxError                        = "42601"
	psqlInsufficientPrivilege              = "42501"
	psqlCannotCoerce                       = "42846"
	psqlGroupingError                      = "42803"
	psqlWindowingError                     = "42P20"
	psqlInvalidRecursion                   = "42P19"
	psqlInvalidForeignKey                  = "42830"
	psqlInvalidName                        = "42602"
	psqlNameTooLong                        = "42622"
	psqlReservedName                       = "42939"
	psqlDatatypeMismatch                   = "42804"
	psqlIndeterminateDatatype              = "42P18"
	psqlCollationMismatch                  = "42P21"
	psqlIndeterminateCollation             = "42P22"
	psqlWrongObjectType                    = "42809"
	psqlUndefinedColumn                    = "42703"
	psqlUndefinedFunction                  = "42883"
	psqlUndefinedTable                     = "42P01"
	psqlUndefinedParameter                 = "42P02"
	psqlUndefinedObject                    = "42704"
	psqlDuplicateColumn                    = "42701"
	psqlDuplicateCursor                    = "42P03"
	psqlDuplicateDatabase                  = "42P04"
	psqlDuplicateFunction                  = "42723"
	psqlDuplicatePreparedStatement         = "42P05"
	psqlDuplicateSchema                    = "42P06"
	psqlDuplicateTable                     = "42P07"
	psqlDuplicateAlias                     = "42712"
	psqlDuplicateObject                    = "42710"
	psqlAmbiguousColumn                    = "42702"
	psqlAmbiguousFunction                  = "42725"
	psqlAmbiguousParameter                 = "42P08"
	psqlAmbiguousAlias                     = "42P09"
	psqlInvalidColumnReference             = "42P10"
	psqlInvalidColumnDefinition            = "42611"
	psqlInvalidCursorDefinition            = "42P11"
	psqlInvalidDatabaseDefinition          = "42P12"
	psqlInvalidFunctionDefinition          = "42P13"
	psqlInvalidPreparedStatementDefinition = "42P14"
	psqlInvalidSchemaDefinition            = "42P15"
	psqlInvalidTableDefinition             = "42P16"
	psqlInvalidObjectDefinition            = "42P17"
	// Class 44 - WITH CHECK OPTION Violation
	psqlWithCheckOptionViolation = "44000"
	// Class 53 - Insufficient Resources
	psqlInsufficientResources      = "53000"
	psqlDiskFull                   = "53100"
	psqlOutOfMemory                = "53200"
	psqlTooManyConnections         = "53300"
	psqlConfigurationLimitExceeded = "53400"
	// Class 54 - Program Limit Exceeded
	psqlProgramLimitExceeded = "54000"
	psqlStatementTooComplex  = "54001"
	psqlTooManyColumns       = "54011"
	psqlTooManyArguments     = "54023"
	// Class 55 - Object Not In Prerequisite State
	psqlObjectNotInPrerequisiteState = "55000"
	psqlObjectInUse                  = "55006"
	psqlCantChangeRuntimeParam       = "55P02"
	psqlLockNotAvailable             = "55P03"
	// Class 57 - Operator Intervention
	psqlOperatorIntervention = "57000"
	psqlQueryCanceled        = "57014"
	psqlAdminShutdown        = "57P01"
	psqlCrashShutdown        = "57P02"
	psqlCannotConnectNow     = "57P03"
	psqlDatabaseDropped      = "57P04"
	// Class 58 - System Error (errors external to PostgreSQL itself)
	psqlSystemError   = "58000"
	psqlIoError       = "58030"
	psqlUndefinedFile = "58P01"
	psqlDuplicateFile = "58P02"
	// Class F0 - Configuration File Error
	psqlConfigFileError = "F0000"
	psqlLockFileExists  = "F0001"
	// Class HV - Foreign Data Wrapper Error (SQL/MED)
	psqlFdwError                             = "HV000"
	psqlFdwColumnNameNotFound                = "HV005"
	psqlFdwDynamicParameterValueNeeded       = "HV002"
	psqlFdwFunctionSequenceError             = "HV010"
	psqlFdwInconsistentDescriptorInformation = "HV021"
	psqlFdwInvalidAttributeValue             = "HV024"
	psqlFdwInvalidColumnName                 = "HV007"
	psqlFdwInvalidColumnNumber               = "HV008"
	psqlFdwInvalidDataType                   = "HV004"
	psqlFdwInvalidDataTypeDescriptors        = "HV006"
	psqlFdwInvalidDescriptorFieldIdentifier  = "HV091"
	psqlFdwInvalidHandle                     = "HV00B"
	psqlFdwInvalidOptionIndex                = "HV00C"
	psqlFdwInvalidOptionName                 = "HV00D"
	psqlFdwInvalidStringLengthOrBufferLength = "HV090"
	psqlFdwInvalidStringFormat               = "HV00A"
	psqlFdwInvalidUseOfNullPointer           = "HV009"
	psqlFdwTooManyHandles                    = "HV014"
	psqlFdwOutOfMemory                       = "HV001"
	psqlFdwNoSchemas                         = "HV00P"
	psqlFdwOptionNameNotFound                = "HV00J"
	psqlFdwReplyHandle                       = "HV00K"
	psqlFdwSchemaNotFound                    = "HV00Q"
	psqlFdwTableNotFound                     = "HV00R"
	psqlFdwUnableToCreateExecution           = "HV00L"
	psqlFdwUnableToCreateReply               = "HV00M"
	psqlFdwUnableToEstablishConnection       = "HV00N"
	// Class P0 - PL/pgSQL Error
	psqlPlpgsqlError   = "P0000"
	psqlRaiseException = "P0001"
	psqlNoDataFound    = "P0002"
	psqlTooManyRows    = "P0003"
	// Class XX - Internal Error
	psqlInternalError  = "XX000"
	psqlDataCorrupted  = "XX001"
	psqlIndexCorrupted = "XX002"
)
