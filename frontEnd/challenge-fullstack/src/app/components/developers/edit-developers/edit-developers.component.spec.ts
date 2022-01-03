import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { EditDevelopersComponent } from './edit-developers.component';

describe('EditStudentComponent', () => {
  let component: EditDevelopersComponent;
  let fixture: ComponentFixture<EditDevelopersComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ EditDevelopersComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(EditDevelopersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
