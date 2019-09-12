package diskmerge

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Source 排序输入源
type Source struct {
	file   *os.File
	reader *bufio.Reader
	line   string
}

// 是否还有下一个元素
func (s *Source) hasNext() bool {
	line, _, err := s.reader.ReadLine()
	s.line = strings.TrimSpace(string(line))
	if err != nil {
		return false
	}
	return true
}

// 获取下一个元素
func (s *Source) next() int {
	if s.line == "" {
		if !s.hasNext() {
			panic(s)
		}
	}
	i, _ := strconv.Atoi(s.line)
	return i
}

// Out 排序输出
type Out struct {
	file   *os.File
	writer *bufio.Writer
}

func (o *Out) write(item MemoryItem) {
	o.writer.WriteString(strconv.Itoa(item.Item) + "\n")
}

func (o *Out) close() {
	if o.writer != nil && o.file != nil {
		o.writer.Flush()
		o.file.Close()
	}
}

// MemoryItem 内存元素
type MemoryItem struct {
	Item   int
	Source Source
}

// MemoryItems 内存排序对象
type MemoryItems []*MemoryItem

func (m MemoryItems) Len() int      { return len(m) }
func (m MemoryItems) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

type ByItem struct {
	MemoryItems
}

func (b ByItem) Less(i, j int) bool { return b.MemoryItems[i].Item < b.MemoryItems[j].Item }

// DiskMergeSort 硬盘多路归并排序
func DiskMergeSort(files []string) {

	items := []*MemoryItem{}

	for _, filename := range files {
		file, _ := os.Open(filename)
		reader := bufio.NewReader(file)
		source := Source{file: file, reader: reader}
		item := MemoryItem{Source: source, Item: source.next()}
		items = append(items, &item)
	}

	// sort items
	sort.Sort(ByItem{items})

	filename := "output.txt"
	file, _ := os.Create(filename)
	wr := bufio.NewWriter(file)
	out := Out{file, wr}
	defer file.Close()
	defer out.close()

	for {
		current := items[0]
		out.write(*current)
		if current.Source.hasNext() {
			current.Item = current.Source.next()
		} else {
			// remove first
			items = items[1:]
			if len(items) == 0 {
				break
			}
		}
		sort.Sort(ByItem{items})
	}
}
