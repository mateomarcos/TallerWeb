import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ListextprojectsComponent } from './listextprojects.component';

describe('ListextprojectsComponent', () => {
  let component: ListextprojectsComponent;
  let fixture: ComponentFixture<ListextprojectsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ListextprojectsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ListextprojectsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
