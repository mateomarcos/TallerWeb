import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ListprojectsComponent } from './listprojects.component';

describe('ListprojectsComponent', () => {
  let component: ListprojectsComponent;
  let fixture: ComponentFixture<ListprojectsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ListprojectsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ListprojectsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
