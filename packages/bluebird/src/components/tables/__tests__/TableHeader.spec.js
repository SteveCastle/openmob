import React from 'react';
import { shallow } from 'enzyme';
import TableHeader from '../TableHeader';

it('renders without crashing', () => {
  shallow(<TableHeader />);
});
