import React from 'react';
import { shallow } from 'enzyme';
import SimpleDonationDrive from '../SimpleDonationDrive';

it('renders without crashing', () => {
  shallow(<SimpleDonationDrive />);
});
