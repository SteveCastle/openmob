import React from 'react';
import { shallow } from 'enzyme';
import TableHeaderCell from '../TableHeaderCell';

it('renders without crashing', () => {
  shallow(<TableHeaderCell />);
});
