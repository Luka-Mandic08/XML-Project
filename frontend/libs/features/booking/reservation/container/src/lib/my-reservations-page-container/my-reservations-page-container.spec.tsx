import { render } from '@testing-library/react';

import MyReservationsPageContainer from './my-reservations-page-container';

describe('MyReservationsPageContainer', () => {
  it('should render successfully', () => {
    const { baseElement } = render(<MyReservationsPageContainer />);
    expect(baseElement).toBeTruthy();
  });
});
