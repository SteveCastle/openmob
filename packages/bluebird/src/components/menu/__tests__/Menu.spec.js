import React from 'react';
import { shallow } from 'enzyme';
import Menu from '../Menu';

it('renders without crashing', () => {
  shallow(<Menu />);
});
