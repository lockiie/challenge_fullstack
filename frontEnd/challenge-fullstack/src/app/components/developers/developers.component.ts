import { Student } from '../../shared/student';
import { ApiService } from '../../shared/api.service';
import { Component, ViewChild, OnInit } from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';



@Component({
  selector: 'app-developers',
  templateUrl: './developers.component.html',
  styleUrls: ['./developers.component.css']
})

export class DevelopersComponent implements OnInit {
  StudentData: any = [];
  dataSource!: MatTableDataSource<Student>;
  @ViewChild(MatPaginator, { static: false }) paginator!: MatPaginator;
  displayedColumns: string[] = ['_id', 'student_name', 'student_email', 'section', 'action'];
  uri : string = 'developers';

  constructor(private studentApi: ApiService) {
    this.studentApi.Get(this.uri).subscribe(data => {
      this.StudentData = data;
      this.dataSource = new MatTableDataSource<Student>(this.StudentData);
      setTimeout(() => {
        this.dataSource.paginator = this.paginator;
      }, 0);
    })    
  }

  ngOnInit() { }

  deleteLevels(index: number, e: any){
    // if(window.confirm('Are you sure')) {
    //   const data = this.dataSource.data;
    //   data.splice((this.paginator.pageIndex * this.paginator.pageSize) + index, 1);
    //   this.dataSource.data = data;
    //   this.studentApi.DeleteStudent(e._id).subscribe()
    // }
  }

}