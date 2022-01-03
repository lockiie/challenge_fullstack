import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AddLevelsComponent } from './add-levels.component';

describe('AddStudentComponent', () => {
  let component: AddLevelsComponent;
  let fixture: ComponentFixture<AddLevelsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AddLevelsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AddLevelsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
