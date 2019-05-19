import React from 'react';
import { shallow } from 'enzyme';
import SideBar from '../SideBar';

it('renders without crashing', () => {
  shallow(<SideBar />);
});
