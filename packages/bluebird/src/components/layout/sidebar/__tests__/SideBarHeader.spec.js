import React from 'react';
import { shallow } from 'enzyme';
import SideBarHeader from '../SideBarHeader';

it('renders without crashing', () => {
  shallow(<SideBarHeader />);
});
