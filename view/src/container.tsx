import * as React from 'react';
import { getData } from "./store";
import Radio from "antd/lib/radio";
import { BarChart, BarChartTag } from "./component/chart";
import { useState } from "react";
import { MapYearButton, MapMonthButton, IfRender } from "./component/button";

export default function () {
  const [ year, setYear ] = useState( "all" );
  const [ month, setMonth ] = useState( "" );
  const [ type, setType ] = useState( "date" );
  return (
    <div className="flex-col">
      <div>
        <div className="flex-jcc mt-50">
          <MapYearButton year={ year } onChange={ setYear }/>
        </div>
        <div className="flex-jcc mt-20">
          <Radio.Group value={ type } onChange={ e => setType( e.target.value ) }>
            <Radio value="date">日期</Radio>
            <Radio value="tag">标签</Radio>
          </Radio.Group>
        </div>
        <div className="flex-jcc mt-20">
          <MapMonthButton year={ year } month={ month } onChange={ setMonth }/>
        </div>
      </div>
      <div className="flex-row">
        <IfRender bool={ type === "date" }>
          <BarChart data={ getData( year, month ) }/>
        </IfRender>
        <IfRender bool={ type === "tag" }>
          <BarChartTag expense data={ getData( year, month ) }/>
          <BarChartTag data={ getData( year, month ) }/>
        </IfRender>
      </div>
    </div>
  )
}


