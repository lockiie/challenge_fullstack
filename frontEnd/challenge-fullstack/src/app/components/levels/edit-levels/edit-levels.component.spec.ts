import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { EditLevelsComponent } from './edit-levels.component';

describe('EditStudentComponent', () => {
  let component: EditLevelsComponent;
  let fixture: ComponentFixture<EditLevelsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ EditLevelsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(EditLevelsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
