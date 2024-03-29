# 진척
- 20230410_x2_error_type.go
  23.2 => page491

</br>

# GO

|제목|내용|
|------|---|
|환경설정|https://learn.microsoft.com/ko-kr/azure/developer/go/configure-visual-studio-code|
|RUN|Ctrl+Alt+N|
|마크다운|https://yunwoong.tistory.com/20|
|패키지설명|https://qiita.com/fetaro/items/31b02b940ce9ec579baf|
|GIT|https://github.com/macmoo/go|
|||

</br>

# 표준 함수
|pkg|함수명|설명|
|---|---|---|
|builtin  |make()|첫 번째 인수로 만들고자 하는 타입을 적어줍니다. 두 번째 인수로 길이|
|         |func copy(dst, src []Type) int|복사|
|bufio    |rd := bufio.NewReader(file)|bufio.Reader 객체생성|
|         |scanner := bufio.NewScanner(file)|스캐너 생성|
|         |line := scanner.Text()|한줄씩 read|
|         |line, _ := rd.ReadString('\n')|bufio.Reader 객체는 구분자까지 문자열을 읽어오는 ReadString() 메서드를 가짐|
|os       |os.Create()|파일생성|
|         |func Open(name string) (*File, error)|파일을 읽기 전용으로 열고 *File 타입인 파일 핸들 객체를 반환|
|fmt      |fmt.Fprintln()||
|         |Scan()|표준 입력에서 값을 입력받습니다.|
|         |Scanf()|표준 입력에서 서식 형태로 값을 입력받습니다.|
|         |Scanln()|표준 입력에서 한 줄을 읽어서 값을 입력받습니다.|
|rand     |func Intn(n int) int|랜덤한 숫자|
|         |func Seed(seed int64)|랜덤 시드 설정|
|         |rand.Seed(time.Now().UnixNano())|랜덤시드설정|
|time     |time.Tick()|일정 시간 간격 주기로 신호를 보내주는 채널을 생성해서 반환하는 함수|
|         |time.After()|현재 시간 이후로 일정 시간 경과 후에 신호를 보내주는 채널을 생성해서 반환하는 함수|
|         |func Now() Time|현재 시각|
|         |func (t Time) UnixNano() int64|int64값으로 변환|
|sync     |var wg sync.WaitGroup||
|         |wg.Add(3)||
|         |wg.Wait()||
|         |wg.Done()||
|context  |context.Background()||
|         |context.WithCancel(ctx) ||
|         |context.WithTimeout(ctx, 시간)||
|         |ontext.WithValue(Ctx, "number", 9)||
|filepath |func Glob(pattern string) (matches []string, err error)|파일 경로에 해당하는 파일 목록을 가져옴|
|         |filespaths, err := filepath.Glob("*.txt")|현재 실행 폴더 내 확장자가 txt인 모든 파일 리스트를 반환|
|strings  |strings.Contains(s, substr string) bool|첫 번째 인수인 s 안에 두 번째 인수인 substr이 포함되어 있는지 여부를 반환하는 함수|
|net/http |func ListenAndServe(addr string, handler Handler) error|웹 서버를 시작|
||||
||||
||||
||||

</br>

# 요약

- fmt 패키지를 이용해서 데이터를 표준 입출력을 할 수 있습니다.
- 표준 출력 함수로는 Print(), Printf(), Println()이 있습니다.
- 서식 문자를 이용하면 다양한 형식으로 출력할 수 있습니다. 최소 출력 너비와 소숫점 이하 숫자 개수를 지정할 수 있습니다.
- 서식 문자 %v를 사용하면 모든 타입의 기본 서식으로 출력합니다.
- 표준 입력 함수로는 Scan(), Scanf(), Scanln()이 있습니다.
- 입력받을 때 에러가 발생하면 표준 입력 스트림을 지웁시다
- 실수 타입은 서로 값이 같은지 비교하는 == 연산자가 비정상 동작할 수 있습니다.
- Go 언어에서 모든 연산자의 각 항의 타입은 항상 같아야 합니다. 대입 연산자도 마찬가지입니다.
- :=(선언대입문) : var 키워드와 타입생략
- Go 내부에서는 필드 각각이 아닌 구조체 전체를 한 번에 복사합니다. </br>
  대입 연산자가 우변 값을 좌변 메모리 공간에 복사할 때 ‘복사되는 크기’는 ‘타입 크기’와 같습니다.</br>
  구조체 크기는 모든 필드를 포함하므로 구조체 전체 필드가 복사되는 겁니다.
- int 크기는 8바이트이고 int32는 4바이트입니다
- unsafe.Sizeof() 함수는 해당 변수의 메모리 공간 크기를 반환
- 메모리 정렬이란 컴퓨터가 데이터에 효과적으로 접근하고자 메모리를 일정 크기 간격으로 정렬하는 것을 말합니다
- 불필요한 메모리 낭비를 줄이려면 작은 크기 필드값을 앞에 배치합시다.
- 메모리 주솟값은 %p로 출력
- 포인터 변숫값을 초기화하지 않으면 기본값은 nil.</br>
  이 값은 0이지만 정확한 의미는 유효하지 않는 메모리 주솟값 즉 어떤 메모리 공간도 가리키고 있지 않음을 나타냄.
- 64비트 컴퓨터에서 메모리 주소는 8바이트이고 32비트 컴퓨터에서는 4바이트 크기를 갖습니다
- 인스턴스란 메모리에 할당된(존재하는) 데이터의 실체를 말합니다.
- Go 언어는 어떤 타입이나 메모리 할당 함수에 의해서 스택 메모리를 사용할지 힙 메모리를 사용할지를 결정하는 게 아닙니다.</br>
  메모리 공간이 함수 외부로 공개되는지 여부를 자동으로 검사해서 스택 메모리에 할당할지 힙 메모리에 할당할지 결정합니다.
- 또 Go 언어에서 스택 메모리는 계속 증가되는 동적 메모리 풀입니다.</br>
  일정한 크기를 갖는 C/C++ 언어와 비교해 메모리 효율성이 높고, 재귀 호출 때문에 스택 메모리가 고갈되는 문제도 발생하지 않습니다.
- UTF-8은 한 글자가 1~3바이트 크기이기 때문에 UTF-8 문자값을 가지려면 3바이트가 필요.</br>
  하지만 Go 언어 기본 타입에서 3바이트 정수 타입은 제공되지 않기 때문에 rune타입은 4바이트 정수 타입인 int32 타입의 별칭 타입이다.</br>
  type rune int32 // rune 타입과 int32는 이름만 다를 뿐 같은 타입입니다.
- 한글은 문자당 3바이트를 차지하고 영문은 1바이트를 차지
- rune은 기본적으로 숫자값
- Go는 string타입에서 *reflect.StringHeader 타입으로 변환을 막고 있기 때문에 string 타입 변수를</br>
  reflect.StringHeader 타입으로 강제 변환을 하려고 unsafe.Pointer(&str1)를 사용해서</br>
  unsafe.Pointer 타입으로 변환한 다음에 다시 *reflect.StringHeader 타입으로 변환합니다.</br>
  내부 필드값을 접근하고자 강제로 타입 변환한다고 이해하면 됩니다.
- 문자열은 불변immutable이다
- string 합 연산을 빈번하게 하면 메모리가 낭비됩니다. 그래서 string 합 연산을 빈번하게 사용하는 경우에는</br>
  strings 패키지의 Builder를 이용해서 메모리 낭비를 줄일 수 있습니다.
- Go 모듈은 Go 패키지들을 모아놓은 Go 프로젝트 단위</br>
  이전까지 Go 모듈을 만들지 않는 Go 코드는 모두 GOPATH/src 폴더</br>
  아래 있어야 했지만 Go 모듈이 기본이 되면서 모든 Go 코드는 Go 모듈 아래 있어야 합니다.</br>
  go build를 하려면 반드시 Go 모듈 루트 폴더에 go.mod 파일이 있어야 합니다</br>
  go.mod 파일은 모듈 이름과 Go 버전, 필요한 외부 패키지 등이 명시되어 있습니다.</br>
  Go 언어에서는 go build를 통해 실행 파일을 만들 때 go.mod와 외부 저장소 패키지 버전 정보를 담고 있는</br>
  go.sum 파일을 통해 외부 패키지와 모듈 내 패키지를 합쳐서 실행 파일을 만들게 됩니다.</br>
  GO모듈생성 => go mod init [패키지명]</br>
  go mod tidy 명령은 Go 모듈에 필요한 패키지를 찾아서 다운로드해주고 필요한 패키지 정보를</br>
  go.mod 파일과 go.sum 파일에 적어준다.
- <r>※ 메모리정렬</r></br>
  레지스터는 실제 연산에 사용되는 데이터가 저장되는 곳.</br>
  레지스터 크기가 4바이트인 컴퓨터를 32비트 컴퓨터라 부르고</br>
  레지스터 크기가 8바이트인 컴퓨터를 64비트 컴퓨터라고 함.(한 번에 8바이트 크기를 연산가능)</br>
  64비트 컴퓨터에서 int64 데이터의 시작 주소가 100번지일 경우 </br>
  100은 8의 배수가 아니기 때문에 레지스터 크기 8에 맞게 정렬되어 있지 않습니다.</br>
  이럴 경우 데이터를 메모리에서 읽어올 때 성능을 손해보기 때문에</br>
  처음부터 프로그램 언어에서 데이터를 만들 때 8의 배수인 메모리 주소에 데이터를 할당한다.</br>
  이 경우 100번지가 아니라 8의 배수인 104번지에 할당.</br>
  메모리 정렬을 위해서 필드 사이에 공간을 띄우는 것을 메모리 패딩Memory Padding이라고 합니다</br>
  참고로 4바이트 변수의 시작 주소는 4의 배수로 맞추고 2바이트 변수의 시작 주소는 2의 배수로 맞춰서 패딩
- <r>배열</r>은 처음 배열을 만들 때 정한 길이에서 더 이상 늘어나지 않는다.</br>
  슬라이스는 배열과 비슷하지만 [ ] 안에 배열의 개수를 적지 않고 선언합니다.</br>
  배열 : var array [10]int</br>
  슬라이스 : var slice []int
- Go 언어에서는 모든 값의 대입은 복사로 일어남. 함수에 인수로 전달될 때나 다른 변수에 대입할 때나 값의 이동은 복사로 일어난다.</br>
  복사는 타입의 값이 복사됩니다. 포인터는 포인터의 값인 메모리 주소가 복사되고 구조체가 복사될 때는 구조체의 모든 필드가 복사된다.</br>
  배열은 배열의 모든 값이 복사된다.
- slice 타입은 []int입니다. []int 타입의 내부는 포인터, len, cap 세 개의 필드를 갖는 구조체입니다.</br>
  포인터는 메모리 주소로 8바이트 길이이고, len과 cap은 각각 int 타입으로 역시 8바이트 길이이기 때문에 슬라이스 크기는 24바이트.</br>
  인수로 slice가 입력되어 호출되면 slice 내부의 포인터가 가리키는 배열 크기에 상관없이 항상 총 24바이트 값이 복사됩니다.
- 슬라이스 내부에는 배열을 가리키는 포인터가 있고, append()는 슬라이스가 가리키는 배열에 빈 공간이 충분하다면 추가하고</br>
  그렇지 않다면 더 큰 배열을 만들어서 추가한다
- 슬라이싱:배열의 일부를 집어내는 기능 => array[startIdx:endindex]
  슬라이싱하면 그 결과로 배열 일부를 가리키는 슬라이스를 반환합니다. 즉, 새로운 배열이 만들어지는 게 아니라</br>
  배열의 일부를 포인터로 가리키는 슬라이스를 만들어낼 뿐입니다.
- array를 슬라이싱할 때  저장 공간 크기를 나타내는 cap은 배열의 총길이에서 시작인덱스를 뺀 만큼을 가지게 됩니다.
- slice[ 시작인덱스 : 끝인덱스 : 최대인덱스 ]</br>
  시작인덱스부터 끝인덱스 하나 전까지 집어내고 최대인덱스까지만 배열을 사용
- 메서드를 선언하려면 리시버를 func 키워드와 함수 이름 사이에 소괄호로 명시해야 합니다.
- func 키워드와 함수명 사이에 리시버 타입과 그 리시버 값을 갖는 변수가 있으면 메서드, 없으면 함수
- 구조체에서 필드가 해당 구조체에 속하듯이 메서드는v해당 리시버 타입에 속한다.
- 모든 사용자 정의 타입이 리시버 타입이 될 수 있기 때문에 기본 내장 타입도 별칭 타입으로 변환하여 메서드를 선언할 수 있다.
- 메서드는 바로 기능과 데이터를 묶어주는 역할, 즉 응집도를 높이는 역할을 한다.
- 좋은 프로그래밍이라면 결합도coupling(객체 간 의존관계)를 낮추고 응집도cohesion(모듈 내 요소들의 상호 관련성)를 높여야 합니다
- 리시버는 메서드를 호출하는 주체로써 메서드는 리시버를 통해서만 호출할 수 있습니다.</br>
  따라서 메서드는 리시버에 속한 기능을 표현합니다. 모든 로컬 타입은 리시버가 될 수 있다.
- 값 타입 메서드 호출 시 값이 모두 복사됩니다. 인스턴스가 아닌 값 중심의 메서드를 만들 때 사용한다.</br>
  호출자 인스턴스에 접근할 수 없고 복사되는 양에 따라서 성능상 문제가 될 수 있다.
- 인터페이스란 구현을 포함하지 않는 메서드 집합
- 추상화 계층을 이용해 의존 관계를 끊는 것을 디커플링decoupling이라함. 의존성이 낮을 수록 좋다.
- Go 언어에서는 어떤 타입이 인터페이스를 포함하고 있는지 여부를 결정할 때 덕 타이핑duck typing</br>
  방식을 사용합니다. 덕 타이핑 방식이란 타입 선언 시 인터페이스 구현 여부를 명시적으로 나타낼</br>
  필요 없이 인터페이스에 정의한 메서드 포함 여부만으로 결정하는 방식입니다.
- 인터페이스 변수의 기본값은 유효하지 않은 메모리 주소를 나타내는 nil이다.</br>
  그래서 인터페이스를 사용할 때 항상 인터페이스값이 nil이 아닌지 확인해야 한다.
- 인터페이스뿐만 아니라 nil값을 기본으로 갖는 다른 타입 변수 역시 사용하기 전에 값이 nil인지 확인해야 한다.</br>
  기본값을 nil로 갖는 타입은 포인터, 인터페이스, 함수 타입, 슬라이스, 맵, 채널 등이 있다.
- 인수 타입 앞에 마침표 3개 ...를 찍어서 <r>가변 인수</r>를 표현
- <r>defer 지연 실행</r></br>
  이와 같이 적으면 명령문이 바로 실행되는 게 아닌 해당 함수가 종료되기 직전에 실행되도록 지연된다.</br>
  명령문은 한 줄의 코드로 일반적으로 함수 호출을 사용</br>
  defer는 역순으로 호출됨
- <r>함수 타입</r></br>
  함수 타입은 함수명과 함수 코드 블록을 제외한 함수 정의function signature로 표시</br>
  별칭 : type opFunc func (int, int) int</br>
  &emsp;&emsp;&emsp;func getOperator(op string) opFunc
  함수 정의에서 매개변수명은 적어도 되고 적지 않아도 됨.
- <r>함수 리터럴function literal</r></br>
  이름 없는 함수로 함수명을 적지 않고 함수 타입 변숫값으로 대입되는 함숫값을 의미.</br>
  함수명이 없기 때문에 함수명으로 직접 함수를 호출할 수 없고 함수 타입변수로만 호출됨.</br>
  다른 프로그래밍 언어에서는 익명 함수 또는 람다Lambda라고 부름.
  함수 리터럴은 필요한 변수를 내부 상태로 가질 수 있다. 함수 리터럴 내부에서 사용되는 외부 변수는 자동으로 함수 내부 상태로 저장된다.
  함수 리터럴에서 외부 변수를 내부 상태로 가져올 때 값 복사가 아닌 인스턴스 참조로 가져오게 됩니다. 포인터 형태로 가져온다고 보면 됨.
- <r>캡쳐capture</r></br>
  함수 리터럴 외부 변수를 내부 상태로 가져오는 것을 캡쳐capture라고 한다. 캡쳐는 값 복사가 아닌 참조 형태로 가져오게 되니 주의해야 한다.
- <r>의존성 주입</r>&emsp;[link][link_ij]</br>
  writeHello() 함수 입장에서 생각해보겠습니다. writeHello()는 인수로 Writer 함수 타입을 받습니다.</br>
  writeHello() 함수 입장에서 보면 인수로 온 writer를 호출했을 때 그게 파일에 씌여질지 네트워크로 전송될지, 프린터로 찍힐지 그게 아니면 아무 상관없는 로직이 수행될지 알 수 없다.</br>
  이렇게 외부에서 로직을 주입하는 것을 의존성 주입이라고 합니다.
- <r>dependency injection</r></br>
  dependency injection. 꼭 함수 리터럴로만 할 수 있는 건 아닙니다. 인터페이스를 통해서도 구현할 수 있습니다.</br>
  외부에서 로직을 주입하는 형태로 구현된다는 것이 중요합니다.
- <r>리스트list</r></br>
  배열과 가장 큰 차이점은 배열은 연속된 메모리에 데이터를 저장하는 반면, 리스트는 불연속된 메모리에 데이터를 저장한다는 점.
- <r>데이터 지역성</r></br>
  데이터 지역성data locality은 데이터가 밀집한 정도를 말합니다. 데이터 로컬리티라고도 한다.
  배열과 리스트를 선택할 때 데이터 지역성을 고려해야 합니다. 컴퓨터는 연산할 때 읽어온 데이터를 캐시라는 임시 저장소에 보관.
  이때 정확히 필요한 데이터만 가져오는 게 아니라 그 주변 데이터를 같이 가져온다. 그 이유는 보통 연산이 일어난 다음에 높은 확률로 주변 데이터에 대한 연산이 이어지기 때문.
  그래서 필요한 데이터가 인접해 있을수록 처리 속도가 빨라지는데, 이를 데이터 지역성이좋다고 말한다.
  배열은 연속된 메모리로 이뤄진 자료구조이고 리스트는 불연속이기 때문에 배열이 리스트에 비해서 데이터 지역성이 월등하게 좋다.
  그래서 삽입과 삭제가 빈번하면 리스트가 배열보다 좋다고 말하지만, 요소 수가 적으면 데이터 지역성 때문에 오히려 배열이 리스트보다 더 효율적입니다. 삽입 삭제가 빈번할 때요소 수에 따른 배열과 리스트 선택은 컴퓨터 성능과 프로그램 성격에 따라 다르다.
- <r>ring</r></br>
  링은 저장할 개수가 고정되고, 오래된 요소는 지워도 되는 경우에 적합
- <r>map</r></br>
  맵map은 키와 값 형태로 데이터를 저장하는 자료구조. 언어에 따라 딕셔너리dictionary, 해시테이블hash table, 해시맵hash map 등으로 부른다. Go 언어에서는 맵이라고 부른다.<br>
  맵은 리스트나 링과 달리 container 패키지가 아닌 Go 기본 내장 타입</br>
  생성 : map[key]value  => map[키타입]값타입</br>
  삭제 : delete(m, key) => delete(맵변수, 삭제키)</br>
  값을 변경하는 방법은 값을 추가할 때와 같습니다. 키에 해당하는 요소가 없으면 추가되고 이미 요소가 있으면 값이 변경됨.</br>
  맵은 내부에서 요소를 보관할 때 입력한 순서와도 키 값과도 상관 없이 데이터를 보관함.</br>
  맵은 반환값을 하나 혹은 둘로 받을 수 있다.</br>
  반환값을 하나만 받으면 값을 반환하고, 둘로 받으면 값뿐 아니라 요소가 존재하는지 알려주는 불리언도 반환.</br>
  반환 : v, ok := m[3] => 값, 존재여부 := m[3]
- <r>해시 함수</r></br>
  해시 함수는 요소 개수와 상관없이 고정된 시간을 갖는 함수이기 때문에 해시 함수를 사용하는 맵이 읽기, 쓰기에서 O(1)의 시간값을 갖게됨.</br>
  또 키가 크다고 해시 함수 결괏값이 커지는 게 아니기 때문에 맵은 키와 무관하고 입력 순서와도 무관한 순서로 순회하게 됨.
- 배열은 연속된 메모리를 사용하고 리스트는 불연속 메모리를 사용.</br>
  리스트는 요소 추가와 삭제 속도가 O(1)입니다. 배열에 비해서 빠르다.
- <r>고루틴</r></br>
  고루틴goroutine은 Go 언어에서 관리하는 경량 스레드</br>
  함수나 명령을 동시에 수행할 때 사용. 여러 고루틴을 갖는 프로그램을코딩하는 것을 동시성프로그래밍Concurrent Programming이라고 한다.</br>
  스레드의 명령 포인터instruction pointer, 스택 메모리 등의 정보를저장하게 되는데 이를 스레드 컨텍스트thread context라고 합니다.</br>
  Go 언어에서는 CPU 코어마다 OS 스레드를 하나만 할당해서 사용하기 때문에컨텍스트 스위칭 비용이 발생하지 않는다.
  고루틴 생성 => go 함수_호출</br>
  go 키워드를 쓰고 함수를 호출하면 해당 함수를 수행하는 새로운 고루틴을생성.</br>호출된함수는 현재 고루틴이 아니라 새로운 고루틴에서 수행된다.
- <r>시스템 콜</r></br>
  시스템 콜이란 운영체제가 지원하는 서비스를 호출할 때를 말합니다. 대표적으로 네트워크 기능등이 있습니다. 시스템 콜을 호출하면 운영체제에서 해당 서비스가 완료될 때까지 대기해야 합니다. 예를 들어 네트워크로 데이터를 읽을 때는 데이터가 들어올 때까지 대기 상태가 됩니다.</br>
  컨텍스트 스위칭은 CPU 코어가 스레드를 변경할 때 발생하는데 고루틴을 이용하면 코어와 스레드는 변경되지 않고 오직 고루틴만 옮겨 다니기 때문입니다. 즉, 코어가 스레드를 변경하지 않기 때문에 컨텍스트 스위칭 비용이 발생하지 않습니다.
- <r>뮤텍스</r></br>
  한 번 획득한 뮤텍스는 반드시 Unlock()을 호출해서 반납해야 한다.</br>
  뮤텍스는 동시에 고루틴 하나만 확보할 수 있습니다
- <r>뮤텍스의 문제점</r></br>
  첫 번째 문제는 동시성 프로그래밍으로 얻는 성능 향상을 얻을 수 없다는 점입니다. 뮤텍스는 오직 하나의 고루틴만 공유 자원에 접근할 수 있도록 제한합니다. 따라서 여러 고루틴 중 뮤텍스를획득한 고루틴만 실행됩니다. 성능을 향상시키려고 동시성 프로그램을 구현했지만 성능 향상을 얻지 못하는 아이러니한 문제가 생기게 됩니다.</br>
  두 번째 문제는 데드락이 발생할 수 있다는 점입니다.
- <r>채널</r></br>
  채널channel과 컨텍스트context는 Go 언어에서 동시성 프로그래밍을 도와주는 기능입니다. 채널은 고루틴 간 메시지를 전달하는 메시지 큐입니다. 채널을 사용하면 뮤텍스 없이 동시성 프로그래밍이 가능합니다. 컨텍스트는 고루틴에 작업을요청할 때 작업 취소나 작업 시간 등을 설정할 수 있는 작업 명세서 역할을 합니다. 채널과 컨텍스트를 사용해 특정 데이터를 전달하거나 특정 시간 동안만 작업을 요청하거나 작업 도중에 작업 취소를 요청할 수 있습니다.</br>
  채널channel이란 고루틴끼리 메시지를 전달할 수 있는 메시지 큐.
  생성:</br>
  var messages chan string = make(chan string)</br>
  var 채널인스턴스변수 채널타입 = make(채널키워드 메시지타입)</br>
  일반적으로 채널을 생성하면 크기가 0인 채널1이 만들어진다.</br>
- 고루틴에서 뮤텍스를 사용하지 않는 방법</br>
  첫 번째 방법인 영역을 나누는 방법</br>
  두 번째 방법인 채널을 이용해서 역할을 나눔.
- <r>생산자 소비자 패턴</r></br>
  데이터를 생성해서 넣어주면 다른 쪽에서 생성된 데이터를 빼서 사용하는 방식을 생산자 소비자 패턴Producer Consumer Pattern이라고 합니다
- <r>컨텍스트</r></br>
  컨텍스트context는 context 패키지에서 제공하는 기능으로 작업을 지시할 때 작업 가능 시간, 작업 취소 등의 조건을 지시할 수 있는 작업 명세서 역할을 합니다
- <r>핸들러</r></br>
  핸들러란 각 HTTP 요청 URL이 수신됐을 때 그것을 처리하는 함수 또는 객체</br>
- <r></r></br>
- 

</br>

# Big-O

|행위|배열, 슬라이스|리스트|
|------|---|---|
|요소 삽입|O(N)|O(1)|
|요소 삭제|O(N)|O(1)|
인덱스 요소 접근|O(1)|O(N)

- 인덱스를 활용한 접근에서는 배열이, 인덱스를 사용한 접근이 거의 없고 삽입과 삭제가 빈번하게 일어나면 리스트가 더 빠릅니다.
- 출력값이 맨 앞에서 발생하기 때문에 배열로 만들면 요소를 빼낼 때마다 O(N) 성능이 필요.</br>
  반면 리스트로 만들면 O(1) 성능을 보장하기 때문에 더 빠르게 처리할 수 있어서 리스트가 큐를 만들 때 더 효율적이다.

</br>

# 컬렉션 속도

맵은 속도가 매우 빠릅니다. 삭제, 추가, 읽기에서 요소 개수와 상관없이 속도가 일정</br>
반면 배열은 추가, 삭제에서 요소 개수가 많아질수록 오래 걸리고,</br>
리스트는 요소 읽기에서 요소 개수가 많아질수록 오래 걸림.</br>
하지만 맵은 키와 값의 쌍으로만 동작하기 때문에 인덱스를 사용해 서 접근할 수 없고 입력한 순서가 보장되지 않는다.</br>
또 배열과 리스트에 비해 상대적으로 메모리를 많이 차지함.</br>

|-|배열, 슬라이스|리스트|맵|
|----|----|----|----|
|추가|O(N)|O(1)|O(1)|
|삭제|O(N)|O(1)|O(1)|
|읽기|O(1) - 인덱스로 접근|O(N) - 인덱스로 접근|O(1) - 키로 접근|




</br>

# 서식지정자

|구분|설명|
|------|---|
|%v|데이터 타입에 맞춰서 기본 형태로 출력합니다.|
|%T|데이터 타입 출력합니다.|
|%t|불리언을 true/false로 출력합니다.|
|%d|10진수 정숫값으로 출력합니다(정수 타입만 가능).|
|%b|2진수로 출력합니다.|
|%c|유니코드 문자를 출력합니다(정수 타입만 가능).|
|%o|8진수로 출력합니다.|
|%O|앞에 8진수임을 표시하는 0o를 붙여서 8진수로 값을 출력합니다.|
|%x|16진수로 값을 출력합니다. 10 이상 값을 a-f 소문자로 표시합니다.|
|%X|16진수로 값을 출력합니다. 10 이상 값을 A-F 대문자로 표시합니다.|
|%e %E|지수 형태로 실숫값을 출력합니다(실수 타입만 가능).예 : -1.234456e+78|
|%f %F|지수 형태가 아닌 실숫값 그대로 출력합니다(실수 타입만 가능).예 : 123.456|
|%g %G|값이 큰 실숫값은 지수 형태(%e)로 출력하고, 작은 실숫값은 실숫값 그대로(%f) 출력합니다.|
|%s|문자열을 출력합니다.||
|%q|특수 문자 기능을 동작하지 않고 문자열 그대로 출력합니다. fmt.Printf("%q", "hello\tWorld\n")|위와 같이 하면 \t 와 \n 특수 문자가 동작하지 않고 hello\tworld\n이 출력됩니다.|
|%p|메모리 주소값을 출력합니다.|

</br>

# 자료형

|타입|최솟값|최댓값|소수부|
|----|-----|-----|-----|
|float32|1.175494351e-38|3.402823466e38|7자리|
|float64|2.2250738585072014e-308|1.7976931348623158e308|15자리|

</br>

# 연산자

|연산자|설명|
|------|----|
|[ ]|배열의 요소에 접근할 때 사용합니다.|
|.|구조체나 패키지 요소에 접근할 때 사용합니다.|
|&|변수의 메모리 주솟값을 반환합니다.|
|*|포인터 변수가 가리키는 메모리 주소에 접근합니다.|
|...|슬라이스 요소들에 접근하거나 가변 인수를 만들 때 사용합니다.|
|:|배열의 일부분을 집어올 때 사용합니다.|
|<-|채널에서 값을 빼거나 넣을 때 사용합니다.|

</br>

# ■ 메모

### <r>※ 포인터가 유효한 메모리 주소를 가리키는지 검사하는 구문</r>

<pre><code>var p *int
if p != nil {
  // p가 nil이 아니라는 얘기는 p가 유효한 메모리 주소를 가리킨다는 뜻입니다.
}</code></pre>

### <r>※ 포인터변수 초기화</r>

<pre><code>p1 := &Data{}             // 1 &를 사용하는 초기화
var p2 = new(Data)        // 2 new()를 사용하는 초기화
// new() 내장 함수는 인수로 타입을 받습니다. 타입을 메모리에 할당하고 기본값으로 채워 그 주소를 반환.
// 1 방식은 p1 := &Data{ 3, 4 }처럼 사용자 초기화가 가능
// 2 방식은 new를 이용해서 내부 필드값을 원하는 값으로 초기화할 수는 없다.
</code></pre>

### <r>※ string 구조</r>

<pre><code>type StringHeader struct {
     Data   uintptr
     Len    int
}
</code></pre>

### <r>※ package</r>

<pre><code>import (
  "text/template"                     // template 패키지
  htemplate "html/template"           // 별칭 htemplate
  _ "github.com/mattn/go-sqlite3"     // (미사용 패키지)밑줄 _을 이용해서 오류 방지
)
</code></pre>

### <r>※ 슬라이스</r>

<pre><code>// ---------------------------------
// 슬라이스 초기화
// ---------------------------------
// 1 { }를 이용해 초기화하기
// 대괄호 안에 길이를 넣지 않은 것을 주의하세요.
var slice1 = []int{1, 2, 3}
// [1 0 0 0 0 2 0 0 0 0 3]
var slice2 = []int{1, 5:2, 10:3}
// ---------------------------------
// 2 make()를 이용한 초기화
// 길이 3개짜리 int 슬라이스값을 갖는다.기본값 0
var slice  = make([]int, 3)
// slice2는 len: 3 cap: 5인 슬라이스가 만들어집니다. 즉 배열 길이는 5, 요소 개수는 3
var slice2 = make([]int, 3, 5)
// ---------------------------------
// 주의
var array = [...]int{1,2,3} // 배열 선언
var slice = []int{1, 2, 3}  // 슬라이스 선언
// ---------------------------------
// 요소 접근
var slice = make([]int, 3)
slice[1] = 5
// ---------------------------------
// 순회
var slice = []int{1, 2, 3}
for i := 0; i < len(slice); i++ { // 1 각 요소에 10 더하기
   slice[i] += 10
}
for i, v := range slice { // 2 각 요소에 2 곱하기
   slice[i] = v * 2
}
// ---------------------------------
</code></pre>

### <r>※ 슬라이스 내부 구조</r>

<pre><code>type SliceHeader struct {
    Data  uintptr // 실제 배열을 가리키는 포인터
    Len   int     // 요소 개수
    Cap   int     // 실제 배열의 길이 , capacity의 약자
}
</code></pre>

### <r>※ 리시버</r>

<pre><code>// (r Rabbit) 부분이 리시버.
// 리서버 덕분에 info() 메서드가 Rabbit 타입에 속한다는 것을 알 수 있다.
// 이때 구조체 변수(r)는 해당 메서드에서 매개변수처럼 사용된다.
// 리시버로는 모든 로컬 타입들이 가능.
// 로컬 타입이란 해당 패키지 안에서 type 키워드로 선언된 타입들을 말함.
// 그래서, 패키지 내 선언된 구조체, 별칭 타입들이 리시버가 될 수 있다.
func (r Rabbit) info() int {
    return r.width * r.height
}
</code></pre>

### <r>※ 인터페이스</r>

<pre><code>// 인터페이스 선언은 1 type을 쓴 뒤 2 인터페이스명을 쓰고 3 interface 키워드를 씁니다.
// 그런 뒤 중괄호 4 { } 블록 안에 인터페이스에 포함된 메서드 집합을 써줍니다
type DuckInterface interface {
    Fly()
    Walk(distance int) int
}
</code></pre>
### <r>※ 인터페이스 변환</r>

<pre><code>// 1. 구체화된 다른 타입으로 타입 변환하기
//    사용 방법은 인터페이스 변수 뒤에 점 .을 찍고 소괄호 () 안에 변경하려는 타입을 써주면 됩니다.
var a Interface
// 인터페이스 변수 a를 ConcreteType 타입으로 변환하고, ConcreteType 타입 변수 t를 생성하고 대입
t := a .(ConcreteType)
// ---------------------------------
// 2. 다른 인터페이스로 타입 변환하기
//    인터페이스 변환을 통해 구체화된 타입뿐 아니라 다른 인터페이스로 타입 변환할 수 있다.
//    이때는 구체화된 타입으로 변환할 때와는 달리 변경되는 인터페이스가 변경 전 인터페이스를 포함
//    하지 않아도 된다. 하지만 인터페이스가 가리키고 있는 실제 인스턴스가 변환하고자 하는 다른 인터페이스를 포함해야 한다.
var a AInterface = ConcreteType{}
b := a.(BInterface)
// AInterface 변수 a는 BInterface로 타입 변환이 가능.
// 그 이유는 a가 가리키고 있는 ConcreteType 인스턴스는 BInterface도 포함 하고 있기 때문.
// ※ 타입 변환 성공 여부 반환
//    타입 변환 반환값을 두 개의 변수로 받으면 타입 변환 가능 여부를 두 번째 반환값(불리언 타입)으로 알려준다.
//    이때 타입 변환이 불가능하더라도 두 번째 반환값이 false로 반환될 뿐 런 타임 에러는 발생하지 않는다.
var a Interface
t, ok := a.(ConcreteType)
// t  : 타입 변환한 결과
// ok : 변환 성공 여부
</code></pre>

### <r>※ 덕 타이핑</r>

<pre><code>// Go 언어에서는 어떤 타입이 인터페이스를 포함하고 있는지 여부를 결정할 때 덕 타이핑duck typing방식을 사용합니다.
//덕 타이핑 방식이란 타입 선언 시 인터페이스 구현 여부를 명시적으로 나타낼 필요 없이 인터페이스에 정의한 메서드 포함 여부만으로 결정하는 방식.
// 인터페이스 정의
type Stringer interface {
    String() string
}
// Student 타입 선언 시 Stringer 인터페이스 포함 여부를 명시적으로 나타내지 않아도
// String() 메서드를 포함하고 있는 것만으로 Stringer 인터페이스를 사용할 수 있다.
type Student struct {
    ...
}
func (s *Student) String() string {
    ...
}
</code></pre>

</br>

### <r>※ Print()</r>

<pre><code>func Print(args ...interface{}) string { // 1 모든 타입을 받는 가변 인수
    // 2 모든 인수 순회
    for _, arg := range args {
        // 3 인수의 타입에 따른 동작
        switch f := arg.(type) {
        case bool:
            // 4 인터페이스 변환
            val := arg.(bool)
            // val 값 출력 로직 생략
        case float64:
            val := arg.(float64)
            // val 값 출력 로직 생략
        case int:
            val := arg.(int)
            // val 값 출력 로직 생략
            // 다른 타입들도 위와 같이 반복
        }
    }
}
</code></pre>

</br>

### <r>※ 리스트를 구현하는 구조체 코드</r>

<pre><code>type Element struct {     // 1 구조체
    Value   interface{ }  // 2 데이터를 저장하는 필드 (interface{ } 타입이므로 어떤 타입값도 저장가능)
    Next    *Element      // 3 다음 요소의 주소를 저장하는 필드
    Prev    *Element      // 4 이전 요소의 주소를 저장하는 필드
}
</code></pre>

</br>

### <r>※ 객체지향 설계 5가지 원칙 SOLID</r>

- 단일 책임 원칙single responsibility principle, SRP
- 개방-폐쇄 원칙open-closed principle, OCP
- 리스코프 치환 원칙liskov substitution principle, LSP
- 인터페이스 분리 원칙interface segregation principle, ISP
- 의존 관계 역전 원칙dependency inversion principle, DIP

<b>나쁜 설계 : ‘상호 결합도가 매우 높고 응집도가 낮은 설계’</b>
‘상호 결합도가 높다’는 얘기는 모듈이 서로 강하게 결합되어 있어서 떼어낼 수 없다는 뜻입니다. 상호 결합도가 높으면 경직성이 증가되고 그로 인해 한 모듈의 수정이 다른 모듈로 전파되어 예기치 못한 문제가 생기고 코드 재사용성을 낮추게 됩니다. ‘응집도가 낮다’는 얘기는 하나의 모듈이 스스로 자립하지 못한다는 뜻입니다. 즉 하나의 모듈이 스스로 완성되지 못하고 다른 모듈에 의존적인 관계를 가지고 있는 경우입니다.

<b>좋은 설계 : ‘상호 결합도가 낮고 응집도가 높은 설계’</b>
상호 결합도가 낮기 때문에 모듈을 쉽게 떼어내서 다른 곳에 사용할 수 있고 모듈 간독립성이 있기 때문에 한 부분을 변경하더라도 다른 모듈에 문제를 발생시키지 않습니다. 그럼으로써 자연스럽게 모듈 완성도가 높아져서 응집도가 높아집니다.

1. 단일 책임 원칙
   - “모든 객체는 책임을 하나만 져야 한다.”
   - 이점 : 코드 재사용성을 높여줍니다.
2. 개방-폐쇄 원칙
   - “확장에는 열려 있고, 변경에는 닫혀 있다.”
   - 이점 : 상호 결합도를 줄여 새 기능을 추가할 때 기존 구현을 변경하지 않아도 됩니다.
3. 리스코프 치환 원칙
   - “q(x)를 타입 T의 객체 x에 대해 증명할 수 있는 속성이라 하자. 그렇다면 S가 T의 하위 타입이라면 q(y)는 타입 S의 객체 y에 대해 증명할 수 있어야 한다.”
   - 이점 : 예상치 못한 작동을 예방할 수 있습니다.
4. 인터페이스 분리 원칙
   - “클라이언트는 자신이 이용하지 않는 메서드에 의존하지 않아야 한다.”
   - 이점 : 인터페이스를 분리하면 불필요한 메서드들과 의존 관계가 끊어져 더 가볍게 인터페이스를 이용할 수 있습니다.
5. 의존 관계 역전 원칙
   - “상위 계층이 하위 계층에 의존하는 전통적인 의존 관계를 반전(역전)시킴으로써 상위 계층이 하위 계층의 구현으로부터 독립되게 할 수 있다.”
    • 원칙 1 : “상위 모듈은 하위 모듈에 의존해서는 안 된다. 둘 다 추상 모듈에 의존해야 한다.”
    • 원칙 2 : “추상 모듈은 구체화된 모듈에 의존해서는 안 된다. 구체화된 모듈은 추상 모듈에 의존해야 한다.”
   - 이점
    • 구체화된 모듈이 아닌 추상 모듈에 의존함으로써 확장성이 증가합니다.
    • 상호 결합도가 낮아져서 다른 프로그램으로 이식성이 증가합니다.

<pre><code>
</code></pre>

</br>

### <r>※ TEST</r>

- 파일명이 _test.go로 끝나야 합니다
- testing 패키지를 임포트해야 합니다.
- 테스트 코드는 func TestXxxx(t *testing.T) 형태이어야 합니다.
- t.Errorf() 메서드에 테스트 실패 시 실패를 알리고 실패 메시지를 넣을 수 있다.</br>
  testing.T 객체의 Error()와 Fail() 메서드를 이용해서 테스트 실패를 알릴 수 있습니다.</br>
  Error()는 테스트가 실패하면 모든 테스트를 중단하지만,</br>
  Fail()은 테스트가 실패해도 다른 테스트들을 계속 진행합니다.

### stretchr

|메소드명|내용||
|---|---|---|
|Equal()     |expected와 actual 두 값을 비교하여 다를 경우 테스트를 실패하고 메시지를 출력.|func Equal(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool|
|Greater()   |e1이 e2보다 크지 않으면 테스트를 실패하고 메시지를 출력.|func Greater(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool|
|Len()       |object의 항목 개수가 length가 아니면 테스트를 실패하고 메시지를 출력합니다.|func Len(t TestingT, object interface{}, length int, msgAndArgs ...interface{}) bool|
|NotNilf()   |object가 nil이면 테스트를 실패하고 메시지를 출력.|func NotNilf(t TestingT, object interface{}, msg string, args ...interface{}) bool|
|NotEqualf() |expected와 actual이 같으면 테스트를 실패하고 메시지를 출력.|func NotEqualf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) bool|

### stretchr/testify 패키지에서 제공하는 그외 유용한 기능

- mock 패키지 : 모듈의 행동을 가장하는 목업mockup 객체를 제공. 예를 들어 온라인 기능을 테스트할 때 하위 영역인 네트워크 기능까지 모두 테스트하기는 힘듭니다. 네트워크 객체를 가장하는 목업 객체를 만들 때 유용합니다.
- suite 패키지 : 테스트 준비 작업이나 테스트 종료 후 뒤처리 작업을 쉽게 할 수 있도록 도와주는 패키지. 예를 들어 테스트에 특정 파일이 있어야 한다면 테스트 시작 전 임시 파일을 생성하고 테스트 종료 후 생성한 임시 파일을 삭제해주는 작업을 만들 때 유용.

</br>

### <r>※ 벤치마크</r>

- 파일명이 _test.go로 끝나야 한다.
- testing 패키지를 임포트해야 한다.
- 벤치마크 코드는 func BenchmarkXxxx(b *testing.B) 형태이어야 한다.

<pre><code>
</code></pre>

</br>

### <r>※ http.Request 구조체</r>

<pre><code>type Request struct {
    // HTTP 요청 메서드 정보를 가지고 있습니다.
    // "GET", "POST", "PUT", “DELETE”와 같은 메서드값입니다.
    Method string

    // HTTP 요청을 보낸 URL 정보를 담고 있습니다.
    // URL 정보를 이용해서 URL에 포함된 데이터를 쿼리해올 수 있습니다.
    URL *url.URL

    // HTTP 프로토콜 버전 정보입니다.
    // HTTP 1.0인지 2.0인지를 알아올 수 있습니다.
    Proto         string // "HTTP/1.0"
    ProtoMajor    int // 1
    ProtoMinor    int // 0

    // HTTP 요청 헤더 정보입니다.
    // 만약 헤더가 다음과 같다면
    //
    // Host: example.com
    // accept-encoding: gzip, deflate
    // Accept-Language: en-us
    // fOO: Bar
    // foo: two
    //
    // 아래와 같이 맵 형태로 변환되어 저장됩니다.
    //
    // Header = map[string][]string{
    //     "Accept-Encoding": {"gzip, deflate"},
    //     "Accept-Language": {"en-us"},
    //     "Foo": {"Bar", "two"},
    // }
    Header Header
    // HTTP 요청의 실제 데이터를 담고 있는 바디 정보입니다.
    // io.Reader 인터페이스를 통해서 데이터를 읽어올 수 있습니다.
    // io.Reader에 대해서는 A.3절을 참조하세요.

    Body io.ReadCloser
    // 그외…
    // http.Request는 이외에도 다양한 정보를 포함하고 있습니다.
}
</code></pre>


# ■ redis
- key, value의 비정형 데이터를 저장해놓는 nosql입니다.</br>
  RDB는 parser, 전처리기, 옵티마이저를 타고 디스크까지 접근하므로 느리지만</br>
  redis(remote dictionary server)는 서버의 메모리에서 동작하며</br>
  조회 정책을 기본적으로 해시테이블방식을 사용하므로(시간복잡도가 O(1)) 빠릅니다.</br>
  대신 정규화가 불가능해 RDB에 비해 확장에 불리합니다.</br>
  만료시간이 있는 데이터를 저장하거나, 확장될 여지가 없는 데이터셋을 저장할때 주로 사용됩니다.</br>
  (redis는 리모트 서버의 인메모리에서 동작하지만 서버가 내려갔을때 데이터 영속성을 위해</br>
  AOF와 RDB방식을 제공한다는 특징이 있습니다.

</br>

# ■ Redis 참조
|제목|링크|
|---|---|
|Hiredis|http://www.redisgate.com/redis/clients/hiredis_intro.php|
|참고|https://github.com/jaeho310/go-redis-client|
|windows 설치|https://github.com/microsoftarchive/redis|
|환경설정|https://pamyferret.tistory.com/9|
|redis api doc|https://pkg.go.dev/github.com/go-redis/redis/v8|
|redis sample|https://redis.uptrace.dev/guide/go-redis.html#connecting-to-redis-server|
|redis sample|https://qiita.com/momotaro98/items/ed496ba06908b278e103|
|redis sentine|https://redis.uptrace.dev/guide/go-redis-sentinel.html|
|redis test|https://frozenpond.tistory.com/164|
|Redisデータ操作|https://www.wakuwakubank.com/posts/736-server-redis-cli/#index_id6|

</br>

# ■ gRPC

- google 에서 만든 오픈소스로, 원격지의 프로시저를 호출하는 프레임워크 입니다.
- 원격지 프로시저를 수행하는 규칙 및 파라미터 전달을 위한 인터페이스로 protocol buffer 라는 오픈소스를 활용.
- Blocking & Non-Blocking 을 지원합니다.
- HTTP/2 프로토콜을 사용합니다.
- 인증, 로드벨런싱, 트레이싱, 헬스체크 등을 제공합니다.
- 10개 언어에서 지원되는 라이브러리가 있습니다.
- gRPC 사용 및 흐름
  1. 실행하고자 하는 프로시저와, 전달하고자 하는 파라미터 사양을 .proto 파일로 작성합니다.</br>
  2. protoc 를 통해 사용하고자 하는 언어에 맞게 stub 파일을 생성합니다.</br>
     생성된 파일은 각 클라이언트가 참조할 수 있는 언어(.java .c .go 등..) 로써 bean 과 같이 데이터를 엑세스 하거나 핸들링하는 함수가 포함되어 있습니다.</br>
  3. gRPC 에서 각 언어별로 제공하는 SDK 를 제공합니다. 이를 활용해 서버, 클라이언트를 프로그래밍 합니다.</br>
  4. stub 을 활용해 실행될 프로시저를 구현하거나 전달할 파라미터를 생산할수 있습니다.
- .proto 와 stub 파일</br>
  protocol buffer 를 사용하는 이점중 하나는</br>
  .proto 파일로 구조화된 데이터를 작성하기만 한다면 gRPC가 지원하는 어떤 언어에서든 규약에 상관없이 통신이 가능하다는 것.</br>
  작성된 .proto 로부터 언어에 맞는 stub 를 생산하여 참조하게 되면 이후 별도의 사양서를 볼 필요없이 참조한 stub 만으로도 개발이 가능.

</br>

# ■ <r>gRPC 참조</r>

|제목|링크|
|---|---|
|참조|https://etloveguitar.tistory.com/107|
||https://dev.classmethod.jp/articles/golang-grpc-sample-project/|
||https://lejewk.github.io/grpc-go-example/|
||https://devjin-blog.com/golang-grpc-server-1/|
|*|https://github.com/dojinkimm/go-grpc-example|
|||
|||
|||
|||

</br>

<style>
r { color: Red }
o { color: Orange }
g { color: Green }
b { color: Blue }
</style>