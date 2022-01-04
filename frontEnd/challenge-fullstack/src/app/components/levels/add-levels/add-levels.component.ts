import { Router } from '@angular/router';
import { Component, OnInit, ViewChild, NgZone } from '@angular/core';
import { ApiService } from '../../../shared/api.service';
import { FormGroup, FormBuilder, Validators } from "@angular/forms";

export interface Subject {
  name: string;
}

@Component({
  selector: 'app-add_levels',
  templateUrl: './add-levels.component.html',
  styleUrls: ['./add-levels.component.css']
})

export class AddLevelsComponent implements OnInit {
  visible = true;
  selectable = true;
  removable = true;
  addOnBlur = true;
  uri : string = 'levels';
  @ViewChild('resetLevelForm', { static: true }) myNgForm: any;
  levelForm!: FormGroup;

  ngOnInit() {
    this.submitBookForm();
  }

  constructor(
    public fb: FormBuilder,
    private router: Router,
    private ngZone: NgZone,
    private levelApi: ApiService
  ) { }

  submitBookForm() {
    this.levelForm = this.fb.group({
      name: ['', [Validators.required]],
    })
  }

  public handleError = (controlName: string, errorName: string) => {
    return this.levelForm.controls[controlName].hasError(errorName);
  }  

  submitLevelForm() {
    if (this.levelForm.valid) {
      console.log("das")
      this.levelApi.Add(this.uri, this.levelForm.value).subscribe(res => {
        this.ngZone.run(() => this.router.navigateByUrl('/levels/list'))
      });
    }
  }

}