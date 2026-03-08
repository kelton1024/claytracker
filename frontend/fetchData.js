window.CESIUM_BASE_URL = 'node_modules/cesium/Build/Cesium';

import { Viewer, SceneMode, Credit, WebMapServiceImageryProvider, Cartesian3, Color, ScreenSpaceEventHandler, ScreenSpaceEventType } from 'cesium'
import "cesium/Build/Cesium/Widgets/widgets.css"

document.addEventListener("DOMContentLoaded", () => {
    const viewer = new Viewer('cesiumContainer', {
        baseLayerPicker: false,
        fullscreenButton: false,
        homeButton: false,
        infoBox: false,
        navigationHelpButton: false,
        sceneModePicker: false,
        timeline: false,
        selectionIndicator: false,
        animation: false,
        geocoder: false,
        sceneMode: SceneMode.SCENE2D
    });

    const shadedRelief1 = new WebMapServiceImageryProvider({
        url: 'https://basemap.nationalmap.gov/arcgis/rest/services/USGSImageryOnly/MapServer',
        layers: 'USGSImageryOnly',
        credit: new Credit('U.S Geological Survey')
    });

    viewer.imageryLayers.addImageryProvider(shadedRelief1)

    viewer.camera.setView({
        destination: Cartesian3.fromRadians(-1.450954407666368, 0.7033698430897186, 900.1678122997862)
    })

    const stations = [
        {longitude: -1.4509169645241433, latitude: 0.7033471365321567},
        {longitude: -1.4509085144602956, latitude: 0.7033466620145326},
        {longitude: -1.4509039700415078, latitude: 0.7033650457825863},
        {longitude: -1.4509029734417584, latitude: 0.7033703768627884},
        {longitude: -1.4509028715123407, latitude: 0.7033761185916061},
        {longitude: -1.4509027101431982, latitude: 0.7033786825679846},
        {longitude: -1.4509040369561488, latitude: 0.7033839001702642},
        {longitude: -1.4509142150079661, latitude: 0.7033870375273719},
        {longitude: -1.4509202017319507, latitude: 0.7033876899267805},
        {longitude: -1.450926725726036, latitude: 0.703387977750049},
        {longitude: -1.4509354563651795, latitude: 0.7033902995244155},
        {longitude: -1.4509441870043236, latitude: 0.7033915275703613},
        {longitude: -1.4509527449495048, latitude: 0.7033910286767},
        {longitude: -1.4509575995686332, latitude: 0.703391182182443},
        {longitude: -1.4509713958973023, latitude: 0.703391143806007},
        {longitude: -1.4509886269169732, latitude: 0.7033905489712546},
        {longitude: -1.4509925796898604, latitude: 0.7033827393665721},
        {longitude: -1.4509925413134188, latitude: 0.7033791895462691},
        {longitude: -1.4509862092015064, latitude: 0.7033512131245927}
    ]

    let i = 1
    stations.forEach((station)=>{
        viewer.entities.add({
            position: Cartesian3.fromRadians(station.longitude, station.latitude, -1),
            point: {
                pixelSize: 30,
                color: Color.BLACK
            },
            label: {
                text: "Station " + i,
                showBackground: true,
                font: '10pt mono',
                eyeOffset: new Cartesian3(0, 0, -50)
            },
            id: i
        })
        i++;
    })

    const handler = new ScreenSpaceEventHandler(viewer.canvas);
    handler.setInputAction((click)=>{
        const obj = viewer.scene.pick(click.position);

        if(!obj){
            return
        }
        
        const id = obj.id._id - 1
        console.log(id)

        const results = document.getElementById("results")
        const rows = results?.children
        for(let i = 0; i < 4; i++){
            const td = rows[i].children
            for(let j = 0; j < 20; j++){
                if(j != id){
                    continue
                }
                if(td[j].classList.contains("bg-blue-900"))
                    td[j].classList.remove("bg-blue-900")
                else
                    td[j].classList.add("bg-blue-900")
            } 
        }
    }, ScreenSpaceEventType.LEFT_CLICK)

});

// TODO: Move this to an edit file
const scores = document.getElementById("results")
let currentSelected = null;
let prevValue = null;

function edit(event){
    if(event.target && event.target.nodeName == "TD"){
        currentSelected = event.target
        prevValue = event.target.innerHTML
        event.target.innerHTML = '<input id="focus" class="w-20">'
        const inputField = document.getElementById("focus")
        inputField.focus();
    }
}

scores.addEventListener("dblclick", edit)
document.addEventListener("click", unfocus)

function unfocus(){
    if(!currentSelected){
        return
    }
    const inputField = document.getElementById("focus")
    if(inputField.value == ""){
        currentSelected.innerHTML = prevValue
    } else {
        currentSelected.innerHTML = inputField.value;
    }
    currentSelected = null;

}