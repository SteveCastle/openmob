import React from 'react';
import { shallow } from 'enzyme';
import TableCell from '../TableCell';

it('renders without crashing', () => {
  shallow(<TableCell />);
});
