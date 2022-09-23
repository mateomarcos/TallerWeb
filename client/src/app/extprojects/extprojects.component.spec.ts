import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ExtprojectsComponent } from './extprojects.component';

describe('ExtprojectsComponent', () => {
  let component: ExtprojectsComponent;
  let fixture: ComponentFixture<ExtprojectsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ExtprojectsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ExtprojectsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
