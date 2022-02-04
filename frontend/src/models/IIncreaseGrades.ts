import { CoursesInterface } from "./ICourse";
import { StudentsInterface } from "./IStudent";
import { GradesInterface } from "./IGrades";

export interface IncreaseGradesInterface {
  ID: number,
  Date: Date,
  Grade_point: number,
  Description: string,
  StudentID: number,
  Student: StudentsInterface,
  GradesID: number,
  Grades: GradesInterface,
  CourseID: number,
  Course: CoursesInterface,
}