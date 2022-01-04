import { Level } from '../../shared/level';
import { ApiService } from '../../shared/api.service';
import { Component, ViewChild, OnInit } from '@angular/core';
import { MatPaginator, PageEvent } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { MatDialog } from '@angular/material/dialog';
import { ConfirmDialogComponent } from '../dialog/confirm/confirm-dialog.component';



@Component({
  selector: 'app-levels',
  templateUrl: './levels.component.html',
  styleUrls: ['./levels.component.css']
})

export class LevelsComponent implements OnInit {
  LevelData: any = [];
  uri: string = 'levels';
  dataSource: MatTableDataSource<Level> = new MatTableDataSource();
  pageEvent!: PageEvent;
  datasource!: null;
  pageSizeOptions: number[] = [5, 10, 12, 15];
  pageIndex!: number;
  pageSize!: number;
  length!: number;
  search_name: any = "";
  @ViewChild(MatPaginator, { static: false }) paginator!: MatPaginator;
  displayedColumns: string[] = ['id', 'name', 'action'];

  constructor(private LevelApi: ApiService,
    public dialog: MatDialog,
  ) {
  }

  public handlePageEvent(event?: PageEvent) {
    this.LevelApi.Get(this.uri, {
      limit: event!.pageSize,
      offset: event!.pageIndex * event!.pageSize
    }).subscribe(data => {
      this.LevelData = data.resultSet;
      this.dataSource = new MatTableDataSource<Level>(this.LevelData);
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


    this.LevelApi.Get(this.uri, params).subscribe(data => {
      this.LevelData = data.resultSet;
      this.dataSource = new MatTableDataSource<Level>(this.LevelData);
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

  deleteLevels(index: number, e: any) {
    const confirmDialog = this.dialog.open(ConfirmDialogComponent, {
      data: {
        title: 'Confirm delete level',
        message: 'Are you sure, you want to delete an level'
      }
    });
    confirmDialog.afterClosed().subscribe(result => {
      if (result === true) {
        this.LevelApi.Delete(this.uri, e.id).subscribe((res => {
          this.loadData();
        }));
      }
    });
  }

}