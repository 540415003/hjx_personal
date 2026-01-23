#!/usr/bin/env python3
"""
批量处理JWT日志，将解析后的数据写入CSV文件
用于压测数据准备

使用方法:
    python process_jwt_to_csv.py                          # 默认处理所有数据
    python process_jwt_to_csv.py --limit 100              # 只处理前100条
    python process_jwt_to_csv.py --start 10 --limit 100   # 从第10条开始，处理100条
    python process_jwt_to_csv.py --append                 # 追加模式（保留现有数据）
"""
import csv
import argparse
import sys
from decode_jwt import (
    extract_jwt_from_log, 
    decode_jwt_payload, 
    extract_fields, 
    COLUMN_ORDER
)

# 默认文件路径
DEFAULT_INPUT_FILE = 'jwt.csv'
DEFAULT_OUTPUT_FILE = 'throttle_test.csv'


def read_jwt_logs(input_file, start=0, limit=None):
    """
    读取JWT日志文件
    
    Args:
        input_file: 输入文件路径
        start: 起始行（0-based，不包含表头）
        limit: 最多处理多少条
    
    Returns:
        日志行列表
    """
    logs = []
    with open(input_file, 'r', encoding='utf-8') as f:
        reader = csv.reader(f)
        # 跳过表头
        next(reader, None)
        
        for idx, row in enumerate(reader):
            if idx < start:
                continue
            if limit is not None and len(logs) >= limit:
                break
            if row and row[0]:
                logs.append(row[0])
    
    return logs


def process_logs(logs, show_progress=True):
    """
    处理日志列表，提取JWT并解析
    
    Args:
        logs: 日志行列表
        show_progress: 是否显示进度
    
    Returns:
        解析后的数据列表
    """
    results = []
    total = len(logs)
    failed = 0
    
    for idx, log_line in enumerate(logs):
        if show_progress and (idx + 1) % 100 == 0:
            print(f"处理进度: {idx + 1}/{total}", end='\r')
        
        # 提取JWT
        jwt_token = extract_jwt_from_log(log_line)
        if not jwt_token:
            failed += 1
            continue
        
        # 解码JWT
        payload = decode_jwt_payload(jwt_token)
        if not payload:
            failed += 1
            continue
        
        # 提取字段
        fields = extract_fields(payload)
        if not fields:
            failed += 1
            continue
        
        results.append(fields)
    
    if show_progress:
        print(f"\n处理完成: 成功 {len(results)} 条, 失败 {failed} 条")
    
    return results


def write_csv(output_file, data, append=False):
    """
    将数据写入CSV文件
    
    Args:
        output_file: 输出文件路径
        data: 数据列表（每项是一个字段字典）
        append: 是否追加模式
    """
    mode = 'a' if append else 'w'
    
    with open(output_file, mode, newline='', encoding='utf-8') as f:
        writer = csv.writer(f)
        
        # 如果是新文件，写入表头
        if not append:
            writer.writerow(COLUMN_ORDER)
        
        # 写入数据
        for fields in data:
            row = [str(fields.get(col, '')) for col in COLUMN_ORDER]
            writer.writerow(row)
    
    print(f"已写入 {len(data)} 条数据到 {output_file}")


def main():
    parser = argparse.ArgumentParser(
        description='批量处理JWT日志，将解析后的数据写入CSV文件'
    )
    parser.add_argument(
        '-i', '--input', 
        default=DEFAULT_INPUT_FILE,
        help=f'输入文件路径 (默认: {DEFAULT_INPUT_FILE})'
    )
    parser.add_argument(
        '-o', '--output', 
        default=DEFAULT_OUTPUT_FILE,
        help=f'输出文件路径 (默认: {DEFAULT_OUTPUT_FILE})'
    )
    parser.add_argument(
        '--start', 
        type=int, 
        default=0,
        help='从第几条开始处理 (0-based，不包含表头)'
    )
    parser.add_argument(
        '--limit', 
        type=int, 
        default=None,
        help='最多处理多少条 (默认: 全部)'
    )
    parser.add_argument(
        '--append', 
        action='store_true',
        help='追加模式，保留输出文件中的现有数据'
    )
    parser.add_argument(
        '--quiet', 
        action='store_true',
        help='静默模式，不显示进度'
    )
    
    args = parser.parse_args()
    
    # 读取日志
    print(f"正在读取 {args.input}...")
    logs = read_jwt_logs(args.input, args.start, args.limit)
    print(f"读取到 {len(logs)} 条日志")
    
    if not logs:
        print("没有找到需要处理的日志")
        return
    
    # 处理日志
    print("正在处理JWT...")
    results = process_logs(logs, show_progress=not args.quiet)
    
    if not results:
        print("没有成功解析的数据")
        return
    
    # 写入CSV
    write_csv(args.output, results, append=args.append)
    
    print("处理完成!")


if __name__ == '__main__':
    main()
