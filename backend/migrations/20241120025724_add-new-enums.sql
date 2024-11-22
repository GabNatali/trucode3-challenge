
-- Create enum type "workclass_enum"
CREATE TYPE "public"."workclass_enum" AS ENUM ('Private', 'Self-emp-not-inc', 'Self-emp-inc', 'Federal-gov', 'Local-gov', 'State-gov', 'Without-pay', 'Never-worked');
-- Create enum type "education_enum"
CREATE TYPE "public"."education_enum" AS ENUM ('Doctorate', '7th-8th','Bachelors', 'Some-college', '11th', 'HS-grad', 'Prof-school', 'Assoc-acdm', 'Assoc-voc', '9th', '7th', '12th', 'Masters', '1st-4th', '10th', '5th-6th', 'Preschool');
-- Create enum type "marital_status_enum"
CREATE TYPE "public"."marital_status_enum" AS ENUM ('Married-civ-spouse', 'Divorced', 'Never-married', 'Separated', 'Widowed', 'Married-spouse-absent', 'Married-AF-spouse');
-- Create enum type "occupation_enum"
CREATE TYPE "public"."occupation_enum" AS ENUM ('Tech-support', 'Craft-repair', 'Other-service', 'Sales', 'Exec-managerial', 'Prof-specialty', 'Handlers-cleaners', 'Machine-op-inspct', 'Adm-clerical', 'Farming-fishing', 'Transport-moving', 'Priv-house-serv', 'Protective-serv', 'Armed-Forces');
-- Create enum type "relationship_enum"
CREATE TYPE "public"."relationship_enum" AS ENUM ('Wife', 'Own-child', 'Husband', 'Not-in-family', 'Other-relative', 'Unmarried');
-- Create enum type "race_enum"
CREATE TYPE "public"."race_enum" AS ENUM ('White', 'Asian-Pac-Islander', 'Amer-Indian-Eskimo', 'Other', 'Black');
-- Create enum type "sex_enum"
CREATE TYPE "public"."sex_enum" AS ENUM ('Male', 'Female');
-- Create enum type "country_enum"
CREATE TYPE "public"."country_enum" AS ENUM ('United-States', 'Cambodia', 'England', 'Puerto-Rico', 'Canada', 'Germany', 'Outlying-US(Guam-USVI-etc)', 'India', 'Japan', 'Greece', 'South', 'China', 'Cuba', 'Iran', 'Honduras', 'Philippines', 'Italy', 'Poland', 'Jamaica', 'Vietnam', 'Mexico', 'Portugal', 'Ireland', 'France', 'Dominican-Republic', 'Laos', 'Ecuador', 'Taiwan', 'Haiti', 'Columbia', 'Hungary', 'Guatemala', 'Nicaragua', 'Scotland', 'Thailand', 'Yugoslavia', 'El-Salvador', 'Trinadad&Tobago', 'Peru', 'Hong', 'Holand-Netherlands');
-- Create enum type "income_enum"
CREATE TYPE "public"."income_enum" AS ENUM ('<=50K', '>50K');
-- Create "adults" table
CREATE TABLE "public"."adults" (
  "age" bigint NULL,
  "workclass" "public"."workclass_enum" NULL,
  "fnlwgt" bigint NULL,
  "education" "public"."education_enum" NULL,
  "education_num" integer NULL,
  "marital_status" "public"."marital_status_enum" NULL,
  "occupation" "public"."occupation_enum" NULL,
  "relationship" "public"."relationship_enum" NULL,
  "race" "public"."race_enum" NULL,
  "sex" "public"."sex_enum" NULL,
  "capital_gain" bigint NULL,
  "capital_loss" bigint NULL,
  "hours_per_week" bigint NULL,
  "native_country" "public"."country_enum" NULL,
  "income" "public"."income_enum" NULL
);
