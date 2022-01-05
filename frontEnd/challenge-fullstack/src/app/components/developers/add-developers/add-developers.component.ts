import { Router } from '@angular/router';
import { Component, OnInit, ViewChild, NgZone } from '@angular/core';
import { COMMA, ENTER } from '@angular/cdk/keycodes';
import { MatChipInputEvent } from '@angular/material/chips';
import { ApiService } from '../../../shared/api.service';
import { FormGroup, FormBuilder, Validators } from "@angular/forms";
import { Level } from 'src/app/shared/level';

export interface Subject {
  name: string;
}

@Component({
  selector: 'app-add_developers',
  templateUrl: './add-developers.component.html',
  styleUrls: ['./add-developers.component.css']
})

export class AddDevelopersComponent implements OnInit {
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
    private Api: ApiService
  ) { }

  ngOnInit() {
    this.submitBookForm();
    this.loadLevels();
  }

  submitBookForm() {
    this.developerForm = this.fb.group({
      name: ['', [Validators.required, Validators.maxLength(50)]],
      gender: [null, [Validators.required]],
      birth: [new Date(), [Validators.required]],
      hobby:['', [Validators.required, Validators.maxLength(200)]],
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
    })
  }

  public handleError = (controlName: string, errorName: string) => {
    return this.developerForm.controls[controlName].hasError(errorName);
  }

  submitDeveloperForm() {
    if (this.developerForm.valid) {
      const date = this.developerForm.get('birth')?.value;
      this.developerForm.patchValue({
        birth: new Date(date).toISOString(), 
      });
      this.Api.Add(this.uri, this.developerForm.value).subscribe(res => {
        this.ngZone.run(() => this.router.navigateByUrl('/developers/list'))
      });
    }
  }

  displayLevel(id: any) {
    if (!id) return '';

    let index = this.LevelData.findIndex(level => level.id === id);
    return this.LevelData[index].name;
  }

}