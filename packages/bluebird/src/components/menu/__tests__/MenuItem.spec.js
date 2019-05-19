import React from 'react';
import { shallow } from 'enzyme';
import MenuItem from '../MenuItem';

it('renders without crashing', () => {
  shallow(<MenuItem />);
});
