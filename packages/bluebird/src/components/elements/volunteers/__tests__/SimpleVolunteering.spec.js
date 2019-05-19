import React from 'react';
import { shallow } from 'enzyme';
import SimpleVolunteering from '../SimpleVolunteering';

it('renders without crashing', () => {
  shallow(<SimpleVolunteering />);
});
