import { ApiService } from '../../shared/api.service';
import { Component, ViewChild, OnInit } from '@angular/core';
import { MatPaginator, PageEvent } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { Developer } from 'src/app/shared/developer';
import { MatDialog } from '@angular/material/dialog';
import { ConfirmDialogComponent } from '../dialog/confirm/confirm-dialog.component';



@Component({
  selector: 'app-developers',
  templateUrl: './developers.component.html',
  styleUrls: ['./developers.component.css']
})

export class DevelopersComponent implements OnInit {
  DeveloperData: any = [];
  uri: string = 'developers';
  dataSource: MatTableDataSource<Developer> = new MatTableDataSource();
  pageEvent!: PageEvent;
  datasource!: null;
  pageSizeOptions: number[] = [5, 10, 12, 15];
  pageIndex!: number;
  pageSize!: number;
  length!: number;
  search_name: any = "";
  @ViewChild(MatPaginator, { static: false }) paginator!: MatPaginator;
  displayedColumns: string[] = ['id', 'name', 'gender', 'birth', 'hobby', 'level', 'action'];

  constructor(private DeveloperApi: ApiService,
    public dialog: MatDialog,
  ) {
  }

  public handlePageEvent(event?: PageEvent) {
    this.DeveloperApi.Get(this.uri, {
      limit: event!.pageSize,
      offset: event!.pageIndex * event!.pageSize
    }).subscribe(data => {
      this.DeveloperData = data.resultSet;
      this.dataSource = new MatTableDataSource<Developer>(this.DeveloperData);
      this.pageIndex = 0;
      this.pageSize = 5;
      this.length = data.total;

      // setTimeout(() => {
      //   this.dataSource.paginator = this.paginator;
      // }, 0);
    })
    return event;
  }

  loadData(search_name: string = '') {
    let params = {
      name: search_name,
      limit: this.pageSize,
      offset: 0
    }


    this.DeveloperApi.Get(this.uri, params).subscribe(data => {
      this.DeveloperData = data.resultSet;
      this.dataSource = new MatTableDataSource<Developer>(this.DeveloperData);
      this.pageIndex = 0;
      this.pageSize = 5;
      this.length = data.total;

      // setTimeout(() => {
      //   this.dataSource.paginator = this.paginator;
      // }, 0);
    })
  }

  ngOnInit() {
    this.pageIndex = 0;
    this.pageSize = 5;
    this.loadData();
  }

  deleteDevelopers(index: number, e: any) {
    const confirmDialog = this.dialog.open(ConfirmDialogComponent, {
      data: {
        title: 'Confirm delete developer',
        message: 'Are you sure, you want to delete an developer'
      }
    });
    confirmDialog.afterClosed().subscribe(result => {
      if (result === true) {
        this.DeveloperApi.Delete(this.uri, e.id).subscribe((res => {
          this.loadData();
        }));
      }
    });
  }

  CalculateAge(birth: any): number {
    const date = new Date(birth);
    const timeDiff = Math.abs(Date.now() - date.getTime());
    return Math.floor((timeDiff / (1000 * 3600 * 24)) / 365.25);
  }

}