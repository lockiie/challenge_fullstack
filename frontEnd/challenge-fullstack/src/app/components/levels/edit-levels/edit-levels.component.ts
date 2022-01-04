import { Router, ActivatedRoute } from '@angular/router';
import { Component, OnInit, ViewChild, NgZone } from '@angular/core';
import { COMMA, ENTER } from '@angular/cdk/keycodes';
import { MatChipInputEvent, MatChipList } from '@angular/material/chips';
import { ApiService } from '../../../shared/api.service';
import { FormGroup, FormBuilder, Validators } from "@angular/forms";
import { MatDialog } from '@angular/material/dialog';
import { ConfirmDialogComponent } from '../../dialog/confirm/confirm-dialog.component';

export interface Subject {
  name: string;
}

@Component({
  selector: 'app-edit_levels',
  templateUrl: './edit-levels.component.html',
  styleUrls: ['./edit-levels.component.css']
})

export class EditLevelsComponent implements OnInit {
  visible = true;
  selectable = true;
  removable = true;
  addOnBlur = true;
  uri : string = 'levels';
  @ViewChild('resetLevelForm', { static: true }) myNgForm: any;
  levelForm!: FormGroup;

  ngOnInit() {
    this.updateBookForm();
  }

  constructor(
    public fb: FormBuilder,
    private router: Router,
    private ngZone: NgZone,
    private actRoute: ActivatedRoute,
    private levelApi: ApiService,
    public dialog: MatDialog
  ) { 
    var id = this.actRoute.snapshot.paramMap.get('id');
    this.levelApi.GetByID(this.uri, Number(id)).subscribe(data => {
      this.levelForm = this.fb.group({
        name: [data.name, [Validators.required]],
      })      
    })    
  }

  updateBookForm() {
    this.levelForm = this.fb.group({
      name: ['', [Validators.required]],
    })
  }

  public handleError = (controlName: string, errorName: string) => {
    return this.levelForm.controls[controlName].hasError(errorName);
  } 

  updateLevelForm() {
    const id = this.actRoute.snapshot.paramMap.get('id');
    const confirmDialog = this.dialog.open(ConfirmDialogComponent, {
      data: {
        title: 'Confirm edit level',
        message: 'Are you sure, you want to edit an level'
      }
    });
    confirmDialog.afterClosed().subscribe(result => {
      if (result === true) {
        this.levelApi.Update(this.uri, Number(id), this.levelForm.value).subscribe( res => {
          this.ngZone.run(() => this.router.navigateByUrl('/levels/list'))
        });
      }
    });

  }
  
}
