import { render } from '@testing-library/react';

import Helpers from './helpers';

describe('Helpers', () => {
  it('should render successfully', () => {
    const { baseElement } = render(<Helpers />);
    expect(baseElement).toBeTruthy();
  });
});
