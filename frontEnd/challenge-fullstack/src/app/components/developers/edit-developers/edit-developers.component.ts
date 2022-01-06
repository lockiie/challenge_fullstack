import { ActivatedRoute, Router } from '@angular/router';
import { Component, OnInit, ViewChild, NgZone } from '@angular/core';
import { ApiService } from '../../../shared/api.service';
import { FormGroup, FormBuilder, Validators } from "@angular/forms";
import { Level } from 'src/app/shared/level';
import { MatDialog } from '@angular/material/dialog';
import { ConfirmDialogComponent } from '../../dialog/confirm/confirm-dialog.component';


@Component({
  selector: 'app-edit_developers',
  templateUrl: './edit-developers.component.html',
  styleUrls: ['./edit-developers.component.css']
})

export class EditDevelopersComponent implements OnInit {
  visible = true;
  selectable = true;
  removable = true;
  addOnBlur = true;
  GenderArray = [{
    description: 'Masculine', value: 'M'
  },
  {
    description: 'Feminine', value: 'F'
  }]
  LevelData: Array<Level> = [];
  uri: string = 'developers';
  @ViewChild('resetDeveloperForm', { static: true }) myNgForm: any;
  developerForm!: FormGroup;

  constructor(
    public fb: FormBuilder,
    private router: Router,
    private ngZone: NgZone,
    private Api: ApiService,
    public dialog: MatDialog,
    private actRoute: ActivatedRoute,
  ) {
  }

  ngOnInit() {
    this.loadLevels();
    this.submitBookForm();
  }

  loadData(){
    const id = this.actRoute.snapshot.paramMap.get('id');
    this.Api.GetByID(this.uri, Number(id)).subscribe(data => {
      const birth = new Date(data.birth);
      this.developerForm = this.fb.group({
        name: [data.name, [Validators.required]],
        gender: [data.gender, [Validators.required]],
        birth: [birth.toISOString().split("T")[0], [Validators.required]],
        hobby: [data.hobby, [Validators.required, Validators.maxLength(200)]],
        level: this.fb.group({
          id: [data.level.id, [Validators.required]],
          name: [data.level.name],
        })
      })
    })

  }

  submitBookForm() {
    this.developerForm = this.fb.group({
      name: ['', [Validators.required, Validators.maxLength(50)]],
      gender: [null, [Validators.required]],
      birth: [new Date(), [Validators.required]],
      hobby: ['', [Validators.required, Validators.maxLength(200)]],
      level: this.fb.group({
        id: ['', [Validators.required]],
        name: [''],
      })
    })
  }

  loadLevels(search_name: string = '') {
    let params = {
      name: search_name,
      limit: 5,
      offset: 0
    }
    this.Api.Get('levels', params).subscribe(data => {
      this.LevelData = data.resultSet;
      this.loadData();
    })
  }

  public handleError = (controlName: string, errorName: string) => {
    return this.developerForm.controls[controlName].hasError(errorName);
  }

  updateDeveloperForm() {
    if (this.developerForm.valid) {
      const id = this.actRoute.snapshot.paramMap.get('id');
      const confirmDialog = this.dialog.open(ConfirmDialogComponent, {
        data: {
          title: 'Confirm edit developer',
          message: 'Are you sure, you want to edit an developer'
        }
      });
      confirmDialog.afterClosed().subscribe(result => {
        if (result === true) {
          const date = this.developerForm.get('birth')?.value;
          this.developerForm.patchValue({
            birth: new Date(date).toISOString(),
          });
          this.Api.Update(this.uri, Number(id), this.developerForm.value).subscribe(res => {
            this.ngZone.run(() => this.router.navigateByUrl('/developers/list'))
          });
        }
      });
    }
  }

  displayLevel(id: any) {
    if (!id) return '';

    let index = this.LevelData.findIndex(level => level.id === id);
    return this.LevelData[index].name;
  }

}