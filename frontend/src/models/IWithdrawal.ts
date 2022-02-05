import { RegisCoursesInterface } from "./IRegisCourse";
import { SemestersInterface } from "./ISemester";
import { StudentsInterface } from "./IStudent";
import { TeachersInterface } from "./ITeacher";

export interface WithdrawalsInterface {
    ID: number,
    StudentID: number,
    Student: StudentsInterface,
    RegisCourseID: number,
    RegisCourse: RegisCoursesInterface,
    TeacherID: number,
    Teacher: TeachersInterface,
    SemesterID: number,
    Semester: SemestersInterface,
    YearTime: number,
    RemainCredit: number,
    Reason: string,
    WithdrawalTime: Date,
}