import { Level } from '../../shared/level';
import { ApiService } from '../../shared/api.service';
import { Component, ViewChild, OnInit } from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
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
  dataSource!: MatTableDataSource<Level>;
  @ViewChild(MatPaginator, { static: false }) paginator!: MatPaginator;
  displayedColumns: string[] = ['id', 'name', 'action'];

  constructor(private LevelApi: ApiService,
    public dialog: MatDialog,
  ) {
    this.loadData();
  }

  loadData() {
    this.LevelApi.Get(this.uri).subscribe(data => {
      this.LevelData = data.resultSet;
      this.dataSource = new MatTableDataSource<Level>(this.LevelData);
      setTimeout(() => {
        this.dataSource.paginator = this.paginator;
      }, 0);
    })
  }

  ngOnInit() { }

  deleteLevels(index: number, e: any) {
    const confirmDialog = this.dialog.open(ConfirmDialogComponent, {
      data: {
        title: 'Confirm delete level',
        message: 'Are you sure, you want to delete an level'
      }
    });
    confirmDialog.afterClosed().subscribe(result => {
      if (result === true) {
        this.LevelApi.Delete(this.uri, e.id).subscribe(( res => {
          this.loadData();
        }));
      }
    });
  }

}