import React from 'react';
import { shallow } from 'enzyme';
import SimpleElection from '../SimpleElection';

it('renders without crashing', () => {
  shallow(<SimpleElection />);
});
