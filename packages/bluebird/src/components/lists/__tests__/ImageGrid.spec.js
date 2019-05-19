import React from 'react';
import { shallow } from 'enzyme';
import ImageGrid from '../ImageGrid';

it('renders without crashing', () => {
  shallow(<ImageGrid />);
});
