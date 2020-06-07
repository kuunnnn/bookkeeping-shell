import Radio from "antd/lib/radio";
import { Store } from "../store";
import * as React from "react";

interface MapYearButtonProps {
  year: string;
  onChange: ( year: string ) => void
}

export function MapYearButton ( { year, onChange }: MapYearButtonProps ) {
  return (
    <Radio.Group value={ year } onChange={ ( e ) => onChange( e.target.value ) }>
      <Radio.Button value="all">全部数据</Radio.Button>
      <MapRender list={ Store.children }>
        { ( v ) => (
          <Radio.Button key={ v.name } value={ v.name }>
            { v.name }年
          </Radio.Button>
        ) }
      </MapRender>
    </Radio.Group>
  )
}

interface MapMonthButtonProps {
  year: string;
  month: string;
  onChange: ( month: string ) => void
}

export function MapMonthButton ( props: MapMonthButtonProps ) {
  const { year, month, onChange } = props;
  const data = Store.children.find( v => v.name === year );
  if ( year === "all" || !data ) {
    return null;
  }
  return (
    <Radio.Group value={ month } onChange={ ( e ) => onChange( e.target.value ) }>
      <MapRender list={ data.children }>
        { ( v ) => (
          <Radio.Button key={ v.name } value={ v.name }>
            { v.name }月
          </Radio.Button>
        ) }
      </MapRender>
    </Radio.Group>
  )
}

export function IfRender ( props: { bool: boolean, children: any } ) {
  if ( props.bool ) {
    return props.children;
  }
  return null;
}

function MapRender<T> ( props: { list: T[], children: ( a: T ) => any } ): any {
  return props.list.map( v => props.children( v ) )
}
