import { initLogin } from "./login.js";
import { pixiGame } from "./pixi.js";

const appDiv = document.getElementById('app');
let currentPage = null;

// 페이지 전환 함수
function navigateTo(page) {
    currentPage = page;
    appDiv.innerHTML = ''; // div 초기화

    if (!appDiv) throw new Error("Container element not found");


    switch (page) {
        case 'lobby':
            appDiv.innerHTML = `<h1>로비 화면</h1><p>게임 시작 버튼을 눌러주세요.</p>`;
            break;
        case 'login':
            initLogin(appDiv); // Pixi 로그인 화면 초기화
            break;
        case 'game':
            //initGame(appDiv); // Pixi 게임 실행
            break;
        case 'ranking':
            appDiv.innerHTML = `<h1>랭킹 화면</h1><ul><li>1위: PlayerA</li><li>2위: PlayerB</li></ul>`;
            break;
        case 'pixi':
            pixiGame(appDiv);
            break;
    }
}

// 초기 로드
navigateTo('login');
