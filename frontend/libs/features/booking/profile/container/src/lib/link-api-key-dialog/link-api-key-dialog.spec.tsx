import { render } from '@testing-library/react';

import LinkApiKeyDialog from './link-api-key-dialog';

describe('LinkApiKeyDialog', () => {
  it('should render successfully', () => {
    const { baseElement } = render(<LinkApiKeyDialog />);
    expect(baseElement).toBeTruthy();
  });
});
