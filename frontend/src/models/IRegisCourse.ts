import { CoursesInterface } from "./ICourse";
import { StudentsInterface } from "./IStudent";

export interface RegisCoursesInterface {
    ID: number,
    StudentID: number,
    Student: StudentsInterface,
    CourseID: number,
    Course: CoursesInterface,
  }