import { Store } from '@ngrx/store';
import { NgrxStateAtom } from 'app/ngrx.reducers';
// import { UserPreference } from 'app/services/user-preferences/user-preferences.model';
import { UserPreferencesService } from 'app/services/user-preferences/user-preferences.service';
import { userPreferencesTimeformat } from 'app/services/user-preferences/user-preferences.selector';
import { takeUntil } from 'rxjs/operators';
import { Subject } from 'rxjs';
  // private isDestroyed: Subject < boolean > = new Subject<boolean>();

export class DateTime {
  // Common date formats for use with moment.js -- https://momentjs.com/docs/#/displaying/format/

  // Format for RFC2822 display
  // Wed, 03 Jul 2019 17:08:53 UTC
  public static readonly RFC2822: string = 'ddd, DD MMM YYYY HH:mm:ss [UTC]';

  // Format for testing user setting timezone
  public static readonly DEMO_MODE: string = 'ddd, DD MMM YYYY HH:mm:ss z Z';

  // Format for date display
  // Tue, 24 Sept 2019
  public static readonly CHEF_DATE_TIME: string = 'ddd, DD MMM YYYY';
  // Tue, 24 Sept 2019 UTC
  public static readonly CHEF_DATE_TIME_ZONE: string = 'ddd, DD MMM YYYY z';

  // Format for time display
  // 09:59
  public static readonly CHEF_HOURS_MINS: string = 'HH:mm';

  // Format for filenames of report downloads
  // 2019-09-24-09:59:59
  public static readonly REPORT_DATE_TIME: string = 'YYYY-MM-DD-HHmmss';

  // Format for short date display
  // 25 Oct 2019
  public static readonly CHEF_SHORT_DATE: string = 'DD MMM YYYY';

  // Format for short date display with timezone
  // 25 Oct 2019 UTC
  public static readonly CHEF_SHORT_DATE_ZONE: string = 'DD MMM YYYY z';

  // Format for date labels in event feed graph
  // Tue, 24 Sept
  public static readonly EVENT_GRAPH_DATE_LABEL: string = 'ddd, DD MMM';

  // Format for time labels in event feed table
  // 09:59
  public static readonly EVENT_TABLE_TIME_LABEL: string = 'HH:mm';

  // Format for day labels in event feed table
  // Tuesday
  public static readonly EVENT_TABLE_DAY_LABEL: string = 'dddd';

  // Format for date labels in event feed table
  // 24 September UTC
  public static readonly EVENT_TABLE_DATE_LABEL: string = 'D MMMM [UTC]';

  // Format for navigating to compliance reports with an endtime set
  // 2019-09-24
  public static readonly REPORT_DATE: string = 'YYYY-MM-DD';

  public userSelectedTimeFormat$;
  private isDestroyed: Subject<boolean> = new Subject<boolean>();

  constructor(
    public userPrefsService: UserPreferencesService,
    private store: Store<NgrxStateAtom>
    ) {
      this.userSelectedTimeFormat$ = this.store.select(userPreferencesTimeformat).pipe(
      takeUntil(this.isDestroyed)
    ).subscribe(data => {
      console.log(data);
      return data.value;
    });
  }

}
