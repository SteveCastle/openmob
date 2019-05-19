import React from 'react';
import { shallow } from 'enzyme';
import MenuHeader from '../MenuHeader';

it('renders without crashing', () => {
  shallow(<MenuHeader />);
});
