package utils

import (
	"gorm.io/gorm"
)

func CreateEnums(db *gorm.DB) error {
	enumStatements := []string{
		// workclass_enum
		`DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'workclass_enum') THEN
				CREATE TYPE workclass_enum AS ENUM ('Private', 'Self-emp-inc', 'Self-emp-not-inc', 'Federal-gov', 'Local-gov', 'State-gov', 'Without-pay', 'Never-worked', 'NULL');
			END IF;
		END $$;`,
		// education_enum
		`DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'education_enum') THEN
				CREATE TYPE education_enum AS ENUM ('7th-8th', 'Doctorate', 'Bachelors', 'Some-college', '11th', 'HS-grad', 'Prof-school', 'Assoc-acdm', 'Assoc-voc', '9th', '7th', '12th', 'Masters', '1st-4th', '10th', '5th-6th', 'Preschool', 'NULL');
			END IF;
		END $$;`,
		// marital_status_enum
		`DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'marital_status_enum') THEN
				CREATE TYPE marital_status_enum AS ENUM ('Married-civ-spouse', 'Divorced', 'Never-married', 'Separated', 'Widowed', 'Married-spouse-absent', 'Married-AF-spouse', 'NULL');
			END IF;
		END $$;`,
		// occupation_enum
		`DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'occupation_enum') THEN
				CREATE TYPE occupation_enum AS ENUM ('Tech-support', 'Craft-repair', 'Other-service', 'Sales', 'Exec-managerial', 'Prof-specialty', 'Handlers-cleaners', 'Machine-op-inspct', 'Adm-clerical', 'Farming-fishing', 'Transport-moving', 'Priv-house-serv', 'Protective-serv', 'Armed-Forces', 'NULL');
			END IF;
		END $$;`,
		// relationship_enum
		`DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'relationship_enum') THEN
				CREATE TYPE relationship_enum AS ENUM ('Wife', 'Own-child', 'Husband', 'Not-in-family', 'Other-relative', 'Unmarried', 'NULL');
			END IF;
		END $$;`,
		// race_enum
		`DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'race_enum') THEN
				CREATE TYPE race_enum AS ENUM ('White', 'Black', 'Asian-Pac-Islander', 'Amer-Indian-Eskimo', 'Other', 'NULL');
			END IF;
		END $$;`,
		// sex_enum
		`DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'sex_enum') THEN
				CREATE TYPE sex_enum AS ENUM ('Male', 'Female','NULL');
			END IF;
		END $$;`,
		// country_enum
		`DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'country_enum') THEN
				CREATE TYPE country_enum AS ENUM ('United-States', 'Cambodia', 'England', 'Puerto-Rico', 'Canada', 'Germany', 'Outlying-US(Guam-USVI-etc)', 'India', 'Japan', 'Greece', 'South', 'China', 'Cuba', 'Iran', 'Honduras', 'Philippines', 'Italy', 'Poland', 'Jamaica', 'Vietnam', 'Mexico', 'Portugal', 'Ireland', 'France', 'Dominican-Republic', 'Laos', 'Ecuador', 'Taiwan', 'Haiti', 'Columbia', 'Hungary', 'Guatemala', 'Nicaragua', 'Scotland', 'Thailand', 'Yugoslavia', 'El-Salvador', 'Trinadad&Tobago', 'Peru', 'Hong', 'Holand-Netherlands','NULL');
			END IF;
		END $$;`,
		// income_enum
		`DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'income_enum') THEN
				CREATE TYPE income_enum AS ENUM ('<=50K', '>50K', 'NULL');
			END IF;
		END $$;`,
	}

	for _, stmt := range enumStatements {
		if tx := db.Exec(stmt); tx.Error != nil {
			return tx.Error
		}
	}

	return nil
}
